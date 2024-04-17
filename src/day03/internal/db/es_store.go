package db

import (
	"day03/internal/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"strings"
)

type Store interface {
	GetPlaces(limit int, offset int) ([]models.Place, int, error)
}

type EsStore struct {
	client *elasticsearch.Client
}

func NewEsStore(esClient *elasticsearch.Client) *EsStore {
	return &EsStore{client: esClient}
}

func (es *EsStore) GetPlaces(limit int, offset int) ([]models.Place, int, error) {
	query := fmt.Sprintf(`{ "size" : %d, "from" : %d, "query": { "match_all": {} } }`, limit, offset)

	auth := map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte("umaradri:123123")),
	}

	res, err := es.client.Search(
		es.client.Search.WithIndex("places"),
		es.client.Search.WithBody(strings.NewReader(query)),
		es.client.Search.WithHeader(auth),
	)

	if err != nil {
		return nil, 0, err
	}

	defer res.Body.Close()

	if res.IsError() {
		var errorResponse map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&errorResponse); err != nil {
			return nil, 0, fmt.Errorf("Error parsing the error response body: %s", err)
		}

		return nil, 0, fmt.Errorf("Error get places: %s: %s", res.Status(), errorResponse["error"].(map[string]interface{})["reason"])
	}

	var result map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, 0, fmt.Errorf("Error decoding response body: %s", err)
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})

	places := make([]models.Place, len(hits))

	for i, hit := range hits {

		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		location := source["location"].(map[string]interface{})

		place := models.Place{
			ID:      int(source["id"].(float64)),
			Name:    source["name"].(string),
			Address: source["address"].(string),
			Phone:   source["phone"].(string),
			Location: models.Location{
				Lat: location["lat"].(float64),
				Lon: location["lon"].(float64),
			},
		}

		places[i] = place
	}

	total := int(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))

	return places, total, nil
}
