package app

import (
	"day02/internal/counter"
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

func RunCounter() {
	var countFlags options.CountOptions
	options.SetupCountOptions(&countFlags)

	if !isValidCombination(countFlags) {
		fmt.Println("Error: Invalid combination of flags. Only one of -l, -m, or -w can be true.")
		return
	}

	paths, err := getPathsForCount()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	counter.Count(paths, countFlags)
}

func getPathForFind() (string, error) {
	args := flag.Args()
	if len(args) < 1 {
		return "", errors.New("Path is not specified\nUsage: ./myFind [-f [-ext extension]] [-d] [-sl] path")
	}

	return args[0], nil
}

func getPathsForCount() ([]string, error) {
	args := flag.Args()
	if len(args) < 1 {
		return nil, errors.New("Path is not specified\nUsage: ./myWc [-w] [-l] [-m] path path1...")
	}

	return args, nil
}

func isValidCombination(flags options.CountOptions) bool {
	count := 0
	if flags.Lines {
		count++
	}
	if flags.Characters {
		count++
	}
	if flags.Words {
		count++
	}
	return count <= 1
}
