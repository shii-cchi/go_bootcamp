package app

import (
	"day00/internal/input"
	"day00/internal/statistics"
	"fmt"
)

func Run() {
	var flags input.OutputOptions
	input.SetupOutputOptions(&flags)

	numbers := input.ScanNumbers()

	if len(numbers) == 0 {
		fmt.Println("Error: empty input")
	} else {
		statistics.GetStatistics(numbers, flags)
	}
}
