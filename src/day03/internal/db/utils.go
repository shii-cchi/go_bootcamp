package db

import (
	"encoding/json"
	"log"
	"os"
)

func GetMappingSchema() []byte {
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
