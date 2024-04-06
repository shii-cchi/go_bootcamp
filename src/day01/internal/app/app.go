package app

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func Run() {
	dbFilename := getDbFilename()
	reader, err := NewRecipeReader(dbFilename)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	recipes, err := reader.Read(dbFilename)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(recipes)
}

func getDbFilename() string {
	var dbFilename string

	flag.StringVar(&dbFilename, "f", "", "Specifies the filename of the database")
	flag.Parse()

	return dbFilename
}

type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
}

type Cake struct {
	Name        string       `json:"name" xml:"name"`
	Time        string       `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type RecipeReader interface {
	Read(dbFilename string) ([]Cake, error)
}

func NewRecipeReader(dbFilename string) (RecipeReader, error) {
	extension := filepath.Ext(dbFilename)

	switch extension {
	case ".json":
		return JSONRecipeReader{}, nil
	case ".xml":
		return XMLRecipeReader{}, nil
	default:
		return nil, errors.New("unsupported file extension")
	}
}

type JSONRecipeReader struct{}

func (r JSONRecipeReader) Read(dbFilename string) ([]Cake, error) {
	file, err := os.Open(dbFilename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var recipes struct {
		Cakes []Cake `json:"cake"`
	}

	decoder := json.NewDecoder(file)

	if err = decoder.Decode(&recipes); err != nil {
		return nil, err
	}

	return recipes.Cakes, nil
}

type XMLRecipeReader struct{}

func (r XMLRecipeReader) Read(dbFilename string) ([]Cake, error) {
	file, err := os.Open(dbFilename)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var recipes struct {
		Cakes []Cake `xml:"cake"`
	}

	decoder := xml.NewDecoder(file)

	if err = decoder.Decode(&recipes); err != nil {
		return nil, err
	}

	return recipes.Cakes, nil
}
