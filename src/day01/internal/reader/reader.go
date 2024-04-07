package reader

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"os"
	"path/filepath"
)

type DBReader interface {
	Read(dbFilename string) (Recipes, error)
}

func NewDBReader(dbFilename string) (DBReader, error) {
	extension := filepath.Ext(dbFilename)

	switch extension {
	case ".json":
		return JsonDBReader{}, nil
	case ".xml":
		return XmlDBReader{}, nil
	default:
		return nil, errors.New("unsupported file extension")
	}
}

type JsonDBReader struct{}

func (r JsonDBReader) Read(dbFilename string) (Recipes, error) {
	file, err := os.Open(dbFilename)

	if err != nil {
		return Recipes{}, err
	}

	defer file.Close()

	var recipes Recipes

	decoder := json.NewDecoder(file)

	if err = decoder.Decode(&recipes); err != nil {
		return Recipes{}, err
	}

	return recipes, nil
}

type XmlDBReader struct{}

func (r XmlDBReader) Read(dbFilename string) (Recipes, error) {
	file, err := os.Open(dbFilename)

	if err != nil {
		return Recipes{}, err
	}

	defer file.Close()

	var recipes Recipes

	decoder := xml.NewDecoder(file)

	if err = decoder.Decode(&recipes); err != nil {
		return Recipes{}, err
	}

	return recipes, nil
}
