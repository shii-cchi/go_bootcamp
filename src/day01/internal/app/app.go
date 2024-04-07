package app

import (
	"day01/internal/converter"
	"day01/internal/reader"
	"flag"
	"fmt"
)

func Run() {
	dbFilename := getDbFilename()

	if dbFilename == "" {
		fmt.Println("Error: no filename provided")
		return
	}

	dbReader, err := reader.NewDBReader(dbFilename)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	recipes, err := dbReader.Read(dbFilename)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	convertedRecipes, err := converter.Convert(recipes, dbFilename)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(convertedRecipes))
}

func getDbFilename() string {
	var dbFilename string

	flag.StringVar(&dbFilename, "f", "", "Specifies the filename of the database")
	flag.Parse()

	return dbFilename
}
