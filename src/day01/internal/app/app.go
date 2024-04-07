package app

import (
	"day01/internal/converter"
	"day01/internal/dbcomparer"
	"day01/internal/dbreader"
	"day01/internal/fscomparer"
	"flag"
	"fmt"
	"os"
)

func RunDBReader() {
	dbFilename := getDbFilenameForRead()

	if dbFilename == "" {
		fmt.Println("Error: no filename provided")
		return
	}

	recipes, err := getDataFromDB(dbFilename)

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

func RunDBComparer() {
	oldDBFilename, newDBFilename := getDbFilenamesForCompare()

	if oldDBFilename == "" || newDBFilename == "" {
		fmt.Println("Error: no filenames provided")
		return
	}

	oldData, err := getDataFromDB(oldDBFilename)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	newData, err := getDataFromDB(newDBFilename)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	dbcomparer.Compare(oldData, newData)
}

func RunFSComparer() {
	oldDumpFilename, newDumpFilename := getDumpsFilenames()

	if oldDumpFilename == "" || newDumpFilename == "" {
		fmt.Println("Error: no filenames provided")
		return
	}

	oldDump, err := os.Open(oldDumpFilename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer oldDump.Close()

	newDump, err := os.Open(newDumpFilename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer newDump.Close()

	fscomparer.Compare(oldDump, newDump)
}

func getDbFilenameForRead() string {
	var dbFilename string

	flag.StringVar(&dbFilename, "f", "", "Specifies the filename of the database")
	flag.Parse()

	return dbFilename
}

func getDbFilenamesForCompare() (string, string) {
	var oldDBFilename string
	var newDBFilename string

	flag.StringVar(&oldDBFilename, "old", "", "Specifies the filename of the old database")
	flag.StringVar(&newDBFilename, "new", "", "Specifies the filename of the new database")
	flag.Parse()

	return oldDBFilename, newDBFilename
}

func getDataFromDB(dbFilename string) (dbreader.Recipes, error) {
	dbReader, err := dbreader.NewDBReader(dbFilename)

	if err != nil {
		return dbreader.Recipes{}, err
	}

	recipes, err := dbReader.Read(dbFilename)

	if err != nil {
		return dbreader.Recipes{}, err
	}

	return recipes, nil
}

func getDumpsFilenames() (string, string) {
	var oldDumpFilename string
	var newDumpFilename string

	flag.StringVar(&oldDumpFilename, "old", "", "Specifies the filename of the old filesystem dump")
	flag.StringVar(&newDumpFilename, "new", "", "Specifies the filename of the new filesystem dump")
	flag.Parse()

	return oldDumpFilename, newDumpFilename
}
