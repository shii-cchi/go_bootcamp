package es_utils

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"day03/internal/config"
	"day03/internal/db"
	"day03/internal/models"
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

func CreateIndexAndUploadData(cfg *config.Config) {
	es := MakeNewEsClient(cfg.EsClientAddress)

	mappingsJSON := db.GetMappingSchema(cfg.MappingSchemaFile)

	createIndex(es, mappingsJSON, cfg.UserName, cfg.UserPassword)

	uploadData(es, cfg.DbFile, cfg.UserName, cfg.UserPassword)
}

func MakeNewEsClient(esClientAddress string) *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: []string{esClientAddress},
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

func createIndex(es *elasticsearch.Client, mappingsJSON []byte, userName, userPassword string) {
	req := prepareCreateIndexRequest(mappingsJSON, userName, userPassword)

	res, err := req.Do(context.Background(), es)

	if err != nil {
		log.Fatalf("Error creating the index: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		makeErrorResponse(res)
	}

	log.Printf("Index places created successfully")
}

func prepareCreateIndexRequest(mappingsJSON []byte, userName, userPassword string) *esapi.IndicesCreateRequest {
	req := esapi.IndicesCreateRequest{
		Index: "places",
		Body:  strings.NewReader(string(mappingsJSON)),
	}

	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(userName+":"+userPassword))

	if req.Header == nil {
		req.Header = make(http.Header)
	}

	req.Header.Set("Authorization", auth)

	return &req
}

func makeErrorResponse(res *esapi.Response) {
	var errorResponse map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&errorResponse); err != nil {
		log.Fatalf("Error parsing the error response body: %s", err)
	}

	log.Fatalf("Error creating the index: %s: %s", res.Status(), errorResponse["error"].(map[string]interface{})["reason"])
}

func uploadData(es *elasticsearch.Client, dbFile, userName, userPassword string) {
	file, err := os.Open(dbFile)

	if err != nil {
		log.Fatalf("Error opening %s file: %s", dbFile, err)
	}

	defer file.Close()

	buf := makeRequestBody(file)

	req := prepareBulkRequest(buf, userName, userPassword)

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

func makeRequestBody(file io.Reader) bytes.Buffer {
	scanner := bufio.NewScanner(file)

	scanner.Scan()

	var buf bytes.Buffer

	id := 1

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")

		lat, err := strconv.ParseFloat(fields[4], 64)

		if err != nil {
			log.Fatalf("Error parsing latitude: %s", err)
		}

		lon, err := strconv.ParseFloat(fields[5], 64)

		if err != nil {
			log.Fatalf("Error parsing longitude: %s", err)
		}

		place := models.Place{
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

func prepareBulkRequest(buf bytes.Buffer, userName, userPassword string) *esapi.BulkRequest {
	req := esapi.BulkRequest{
		Body: &buf,
	}

	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(userName+":"+userPassword))

	if req.Header == nil {
		req.Header = make(http.Header)
	}

	req.Header.Set("Authorization", auth)

	return &req
}
