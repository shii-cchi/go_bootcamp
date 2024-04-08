package counter

import (
	"bufio"
	"day02/internal/options"
	"fmt"
	"os"
	"strings"
)

type CountResult struct {
	Path   string
	Result int
	Err    error
}

func Count(paths []string, countFlags options.CountOptions) {
	results := make([]CountResult, len(paths))
	ch := make(chan CountResult)

	for _, path := range paths {
		go func(path string) {
			ch <- countParameters(path, countFlags)
		}(path)
	}

	for i := 0; i < len(paths); i++ {
		results[i] = <-ch
	}

	close(ch)

	for _, result := range results {
		if result.Err != nil {
			fmt.Printf("Error: %v for file %s\n", result.Err, result.Path)
		} else {
			fmt.Println(result.Result, result.Path)
		}
	}
}

func countParameters(path string, countFlags options.CountOptions) CountResult {
	file, err := os.Open(path)

	if err != nil {
		return CountResult{Path: path, Err: err}
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0

	for scanner.Scan() {
		if countFlags.Lines {
			result++
		}

		if countFlags.Characters {
			result += len(scanner.Text())
		}

		if countFlags.Words {
			result += len(strings.Fields(scanner.Text()))
		}
	}

	if err := scanner.Err(); err != nil {
		return CountResult{Path: path, Err: err}
	}

	return CountResult{
		Path:   path,
		Result: result,
	}

}
