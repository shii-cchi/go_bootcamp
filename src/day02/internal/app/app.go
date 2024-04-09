package app

import (
	"day02/internal/archiver"
	"day02/internal/counter"
	"day02/internal/finder"
	"day02/internal/options"
	"day02/internal/runner"
	"errors"
	"flag"
	"fmt"
	"os"
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

	err = finder.Find(path, findFlags)

	if err != nil {
		fmt.Println(err)
		return
	}
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

func RunRunner() {
	command, err := getCommandForRunner()

	if err != nil {
		fmt.Println(err)
		return
	}

	err = runner.Run(command)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func RunArchiver() {
	dir := options.SetupArchiverOptions()

	err := checkDir(dir)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	paths, err := getPathsForArchive()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = archiver.Archive(dir, paths)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
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

func getCommandForRunner() ([]string, error) {
	args := os.Args[1:]
	if len(args) == 0 {
		return nil, errors.New("Error: not enough arguments\nUsage: ./myXargs <command>")
	}

	return args, nil
}

func getPathsForArchive() ([]string, error) {
	args := flag.Args()
	if len(args) < 1 {
		return nil, errors.New("Path is not specified\nUsage: ./myRotate [-a dir]] path path1...")
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

func checkDir(dir string) error {
	dirInfo, err := os.Stat(dir)

	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("Directory does not exist")
		}

		return err
	}

	if !dirInfo.IsDir() {
		return errors.New("Path is not a directory")
	}

	if _, err := os.Stat(dir); err != nil {
		if os.IsPermission(err) {
			return errors.New("Permission denied")
		}
	}

	return nil
}
