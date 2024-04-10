package input

import (
	"day02/internal/options"
	"errors"
	"flag"
	"os"
)

func GetPathForFind() (string, error) {
	args := flag.Args()
	if len(args) < 1 {
		return "", errors.New("Path is not specified\nUsage: ./myFind [-f [-ext extension]] [-d] [-sl] path")
	}

	return args[0], nil
}

func GetPathsForCount() ([]string, error) {
	args := flag.Args()
	if len(args) < 1 {
		return nil, errors.New("Path is not specified\nUsage: ./myWc [-w] [-l] [-m] path path1...")
	}

	return args, nil
}

func GetCommandForRunner() ([]string, error) {
	args := os.Args[1:]
	if len(args) == 0 {
		return nil, errors.New("Error: not enough arguments\nUsage: ./myXargs <command>")
	}

	return args, nil
}

func GetPathsForArchive() ([]string, error) {
	args := flag.Args()
	if len(args) < 1 {
		return nil, errors.New("Path is not specified\nUsage: ./myRotate [-a dir]] path path1...")
	}

	return args, nil
}

func IsValidCombination(flags options.CountOptions) bool {
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

func CheckDir(dir string) error {
	dirInfo, err := os.Stat(dir)

	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("directory does not exist")
		}

		return err
	}

	if !dirInfo.IsDir() {
		return errors.New("path is not a directory")
	}

	if _, err := os.Stat(dir); err != nil {
		if os.IsPermission(err) {
			return errors.New("permission denied")
		}
	}

	return nil
}
