package db

import (
	"day03/internal/es_utils"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"strings"
)

type Store interface {
	GetPlaces(limit int, offset int) ([]Place, int, error)
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
		return nil, 0, fmt.Errorf("Error: %s", es_utils.MakeErrorResponse(res))
	}

	var result map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, 0, fmt.Errorf("Error decoding response body: %s", err)
	}

	places, total := getDataFromResult(result)

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

func getDataFromResult(result map[string]interface{}) ([]Place, int) {
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

	total := int(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	return places, total
}
