package app

import (
	"day00/internal/input"
	"day00/internal/statistics"
	"fmt"
	"os"
)

func Run() {
	var flags input.OutputOptions
	input.SetupOutputOptions(&flags)

	numbers := input.ScanNumbers()

	if len(numbers) == 0 {
		fmt.Fprintln(os.Stderr, "Error: empty input")
	} else {
		statistics.GetStatistics(numbers, flags)
	}
}
