package es_utils

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"day03/internal/db"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func CreateIndexAndUploadData() {
	es := MakeNewEsClient()

	mappingsJSON := getMappingSchema()

	createIndex(es, mappingsJSON)

	uploadData(es)
}

func MakeNewEsClient() *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{"https://localhost:9200"},
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		log.Fatalf("Error creating the Elasticsearch client: %s", err)
	}

	return es
}

func getMappingSchema() []byte {
	file, err := os.Open("schema.json")

	if err != nil {
		log.Fatalf("Error opening schema.json file: %s", err)
	}

	defer file.Close()

	var mappings map[string]interface{}

	if err := json.NewDecoder(file).Decode(&mappings); err != nil {
		log.Fatalf("Error decoding schema.json: %s", err)
	}

	mappingsJSON, err := json.Marshal(mappings)

	if err != nil {
		log.Fatalf("Error marshaling the mappings: %s", err)
	}

	return mappingsJSON
}

func createIndex(es *elasticsearch.Client, mappingsJSON []byte) {
	req := prepareCreateIndexRequest(mappingsJSON)

	res, err := req.Do(context.Background(), es)

	if err != nil {
		log.Fatalf("Error creating the index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", makeErrorResponse(res))
	}

	log.Printf("Index places created successfully")
}

func prepareCreateIndexRequest(mappingsJSON []byte) *esapi.IndicesCreateRequest {
	req := esapi.IndicesCreateRequest{
		Index: "places",
		Body:  strings.NewReader(string(mappingsJSON)),
	}

	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte("umaradri:123123"))

	if req.Header == nil {
		req.Header = make(http.Header)
	}

	req.Header.Set("Authorization", auth)

	return &req
}

func makeErrorResponse(res *esapi.Response) error {
	var errorResponse map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&errorResponse); err != nil {
		return fmt.Errorf("Error parsing the error response body: %s", err)
	}
	return fmt.Errorf("Error: %s: %s", res.Status(), errorResponse["error"].(map[string]interface{})["reason"])
}

func uploadData(es *elasticsearch.Client) {
	file, err := os.Open("data.csv")

	if err != nil {
		log.Fatalf("Error opening data.csv file: %s", err)
	}

	defer file.Close()

	buf := makeBulkRequestBody(file)

	req := prepareBulkRequest(buf)

	res, err := req.Do(context.Background(), es)

	if err != nil {
		log.Fatalf("Error executing Bulk request: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", makeErrorResponse(res))
	}

	log.Println("Bulk request executed successfully")
}

func makeBulkRequestBody(file io.Reader) bytes.Buffer {
	scanner := bufio.NewScanner(file)

	scanner.Scan()

	var buf bytes.Buffer

	id := 1

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")

		lat, err := strconv.ParseFloat(fields[5], 64)

		if err != nil {
			log.Fatalf("Error parsing latitude: %s", err)
		}

		lon, err := strconv.ParseFloat(fields[4], 64)

		if err != nil {
			log.Fatalf("Error parsing longitude: %s", err)
		}

		place := db.Place{
			ID:      id,
			Name:    fields[1],
			Address: fields[2],
			Phone:   fields[3],
			Location: struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			}{Lat: lat, Lon: lon},
		}

		jsonData, err := json.Marshal(place)

		if err != nil {
			log.Fatalf("Error marshaling Place to JSON: %s", err)
		}

		indexLine := []byte(fmt.Sprintf(`{ "index" : { "_index" : "places", "_id" : "%d" } }`+"\n", id))
		buf.Grow(len(indexLine) + len(jsonData) + 1)
		buf.Write(indexLine)
		buf.Write(jsonData)
		buf.WriteByte('\n')

		id++
	}

	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			log.Fatalf("Error reading CSV record: %s", err)
		}
	}

	return buf
}

func prepareBulkRequest(buf bytes.Buffer) *esapi.BulkRequest {
	req := esapi.BulkRequest{
		Body: &buf,
	}

	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte("umaradri:123123"))

	if req.Header == nil {
		req.Header = make(http.Header)
	}

	req.Header.Set("Authorization", auth)

	return &req
}

type CountResponse struct {
	Count int `json:"count"`
}

func GetIndexDocCount(indexName string) (int, error) {
	url := fmt.Sprintf("https://localhost:9200/%s/_count", indexName)

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.SetBasicAuth("umaradri", "123123")

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	var countResponse CountResponse
	err = json.NewDecoder(res.Body).Decode(&countResponse)
	if err != nil {
		return 0, err
	}

	return countResponse.Count, nil
}
