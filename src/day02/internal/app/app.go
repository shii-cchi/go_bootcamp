package app

import (
	"day02/internal/finder"
	"day02/internal/options"
	"errors"
	"flag"
	"fmt"
)

func RunFinder() {
	var findFlags options.FindOptions
	options.SetupFindOptions(&findFlags)

	if findFlags.Extension != "" && !findFlags.File {
		fmt.Println("Error: Flag -ext works only with flag -f")
		return
	}

	path, err := getPathForFind()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	finder.Find(path, findFlags)
}

func getPathForFind() (string, error) {
	args := flag.Args()
	if len(args) != 1 {
		return "", errors.New("Path is not specified/nUsage: ./myFind [-f [-ext extension]] [-d] [-sl] path")
	}

	return args[0], nil
}
