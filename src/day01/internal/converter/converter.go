package converter

import (
	"day01/internal/dbreader"
	"encoding/json"
	"encoding/xml"
	"errors"
	"path/filepath"
)

func Convert(recipes dbreader.Recipes, dbFilename string) ([]byte, error) {
	extension := filepath.Ext(dbFilename)

	switch extension {
	case ".json":
		return convertToXML(recipes)
	case ".xml":
		return convertToJSON(recipes)
	default:
		return nil, errors.New("unsupported file extension")
	}
}

func convertToXML(recipes dbreader.Recipes) ([]byte, error) {
	xmlData, err := xml.MarshalIndent(recipes, "", "  ")

	if err != nil {
		return nil, err
	}

	return xmlData, nil
}

func convertToJSON(recipes dbreader.Recipes) ([]byte, error) {
	jsonData, err := json.MarshalIndent(recipes, "", "  ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
