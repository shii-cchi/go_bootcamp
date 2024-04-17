package db

import (
	"encoding/json"
	"log"
	"os"
)

func GetMappingSchema(mappingSchemaFile string) []byte {
	file, err := os.Open(mappingSchemaFile)

	if err != nil {
		log.Fatalf("Error opening %s file: %s", mappingSchemaFile, err)
	}

	defer file.Close()

	var mappings map[string]interface{}

	if err := json.NewDecoder(file).Decode(&mappings); err != nil {
		log.Fatalf("Error decoding %s: %s", mappingSchemaFile, err)
	}

	mappingsJSON, err := json.Marshal(mappings)

	if err != nil {
		log.Fatalf("Error marshaling the mappings: %s", err)
	}

	return mappingsJSON
}
