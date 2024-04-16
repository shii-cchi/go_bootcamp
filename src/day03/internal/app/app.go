package app

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Place struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Location struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
}

func Run() {

	es := makeNewEsClient()

	//mappingsJSON := getMappingSchema()

	//createIndex(es, mappingsJSON)

	uploadData(es)
}

func makeNewEsClient() *elasticsearch.Client {
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
	req := esapi.IndicesCreateRequest{
		Index: "places",
		Body:  strings.NewReader(string(mappingsJSON)),
	}

	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte("umaradri:123123"))

	if req.Header == nil {
		req.Header = make(http.Header)
	}

	req.Header.Set("Authorization", auth)

	res, err := req.Do(context.Background(), es)

	if err != nil {
		log.Fatalf("4Error creating the index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		var errorResponse map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&errorResponse); err != nil {
			log.Fatalf("Error parsing the error response body: %s", err)
		}

		log.Fatalf("Error creating the index: %s: %s", res.Status(), errorResponse["error"].(map[string]interface{})["reason"])
	}

	log.Printf("Index places created successfully")
}

func uploadData(es *elasticsearch.Client) {
	file, err := os.Open("data.csv")

	if err != nil {
		log.Fatalf("Error opening data.csv file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	var buf bytes.Buffer

	id := 1

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("Error reading CSV record: %s", err)
		}

		lat, err := strconv.ParseFloat(fields[4], 64)
		if err != nil {
			log.Fatalf("Error parsing latitude: %s", err)
		}
		lon, err := strconv.ParseFloat(fields[5], 64)
		if err != nil {
			log.Fatalf("Error parsing longitude: %s", err)
		}

		place := Place{
			ID:      id,
			Name:    fields[1],
			Address: fields[2],
			Phone:   fields[3],
			Location: struct {
				Lat float64 `json:"lat"`
				Lon float64 `json:"lon"`
			}{Lat: lat, Lon: lon},
		}

		id++

		jsonData, err := json.Marshal(place)

		if err != nil {
			log.Fatalf("Error marshaling Place to JSON: %s", err)
		}

		indexLine := []byte(`{ "index" : { "_index" : "places" } }` + "\n")
		buf.Grow(len(indexLine) + len(jsonData) + 1)
		buf.Write(indexLine)
		buf.Write(jsonData)
		buf.WriteByte('\n')
	}

	req := esapi.BulkRequest{
		Body: &buf,
	}
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error executing Bulk request: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error response received from Elasticsearch: %s", res.String())
	}

	log.Println("Bulk request executed successfully")

}
