package db

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"strings"
)

type Store interface {
	GetPlaces(limit int, offset int) ([]Place, int, error)
	GetClosestPlaces(lat, lon float64, limit int) ([]Place, error)
}

type EsStore struct {
	client *elasticsearch.Client
}

func NewEsStore(esClient *elasticsearch.Client) *EsStore {
	return &EsStore{client: esClient}
}

func (es *EsStore) GetPlaces(limit int, offset int) ([]Place, int, error) {
	res, err := executeSearchQuery(es.client, limit, offset)

	if err != nil {
		return nil, 0, err
	}

	defer res.Body.Close()

	if res.IsError() {
		var errorResponse map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&errorResponse); err != nil {
			return nil, 0, fmt.Errorf("Error parsing the error response body: %s", err)
		}

		return nil, 0, fmt.Errorf("Error: %s: %s", res.Status(), errorResponse["error"].(map[string]interface{})["reason"])
	}

	var result map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, 0, fmt.Errorf("Error decoding response body: %s", err)
	}

	places := getDataFromResult(result)

	total := int(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	return places, total, nil
}

func executeSearchQuery(client *elasticsearch.Client, limit, offset int) (*esapi.Response, error) {
	query := fmt.Sprintf(`{ "size" : %d, "from" : %d, "query": { "match_all": {} } }`, limit, offset)

	auth := map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte("umaradri:123123")),
	}

	return client.Search(
		client.Search.WithIndex("places"),
		client.Search.WithBody(strings.NewReader(query)),
		client.Search.WithHeader(auth),
		client.Search.WithTrackTotalHits(true),
	)
}

func getDataFromResult(result map[string]interface{}) []Place {
	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})

	places := make([]Place, len(hits))

	for i, hit := range hits {

		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		location := source["location"].(map[string]interface{})

		place := Place{
			ID:      int(source["id"].(float64)),
			Name:    source["name"].(string),
			Address: source["address"].(string),
			Phone:   source["phone"].(string),
			Location: Location{
				Lat: location["lat"].(float64),
				Lon: location["lon"].(float64),
			},
		}

		places[i] = place
	}

	return places
}

func (es *EsStore) GetClosestPlaces(lat, lon float64, limit int) ([]Place, error) {
	res, err := executeSortQuery(es.client, lat, lon, limit)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.IsError() {
		var errorResponse map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&errorResponse); err != nil {
			return nil, fmt.Errorf("Error parsing the error response body: %s", err)
		}

		return nil, fmt.Errorf("Error: %s: %s", res.Status(), errorResponse["error"].(map[string]interface{})["reason"])
	}

	var result map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Error decoding response body: %s", err)
	}

	places := getDataFromResult(result)

	return places, nil
}

func executeSortQuery(client *elasticsearch.Client, lat, lon float64, limit int) (*esapi.Response, error) {
	query := fmt.Sprintf(`{"size": %d,"sort":[{"_geo_distance":{"location":{"lat": %f,"lon": %f},"order":"asc","unit":"km","mode":"min","distance_type":"arc","ignore_unmapped":true}}]}`, limit, lat, lon)

	auth := map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte("umaradri:123123")),
	}

	return client.Search(
		client.Search.WithIndex("places"),
		client.Search.WithBody(strings.NewReader(query)),
		client.Search.WithHeader(auth),
	)
}
