package app

import (
	"day01/internal/converter"
	"day01/internal/dbcomparer"
	"day01/internal/dbreader"
	"day01/internal/fscomparer"
	"day01/internal/input"
	"fmt"
	"os"
)

func RunDBReader() {
	dbFilename := input.GetDbFilenameForRead()

	if dbFilename == "" {
		fmt.Println("Error: no filename provided")
		return
	}

	recipes, err := dbreader.GetDataFromDB(dbFilename)

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
	oldDBFilename, newDBFilename := input.GetDbFilenamesForCompare()

	if oldDBFilename == "" || newDBFilename == "" {
		fmt.Println("Error: no filenames provided")
		return
	}

	oldData, err := dbreader.GetDataFromDB(oldDBFilename)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	newData, err := dbreader.GetDataFromDB(newDBFilename)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	dbcomparer.Compare(oldData, newData)
}

func RunFSComparer() {
	oldDumpFilename, newDumpFilename := input.GetDumpsFilenames()

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
