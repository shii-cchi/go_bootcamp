package converter

import (
	"day01/internal/reader"
	"encoding/json"
	"encoding/xml"
	"errors"
	"path/filepath"
)

func Convert(recipes reader.Recipes, dbFilename string) ([]byte, error) {
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

func convertToXML(recipes reader.Recipes) ([]byte, error) {
	xmlData, err := xml.MarshalIndent(recipes, "", "  ")

	if err != nil {
		return nil, err
	}

	return xmlData, nil
}

func convertToJSON(recipes reader.Recipes) ([]byte, error) {
	jsonData, err := json.MarshalIndent(recipes, "", "  ")
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
