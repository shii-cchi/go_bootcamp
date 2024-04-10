package statistics

import (
	"day00/internal/input"
	"fmt"
	"math"
	"sort"
)

func GetStatistics(numbers []int, flags input.OutputOptions) {
	if flags.Mean {
		fmt.Printf("Mean: %.2f\n", getMean(numbers))
	}

	if flags.Median {
		fmt.Printf("Median: %.2f\n", getMedian(numbers))
	}

	if flags.Mode {
		fmt.Printf("Mode: %d\n", getMode(numbers))
	}

	if flags.StandardDeviation {
		fmt.Printf("Standard Deviation: %.2f\n", getSD(numbers))
	}
}

func getMean(numbers []int) float64 {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	average := float64(sum) / float64(len(numbers))
	return average
}

func getMedian(numbers []int) float64 {
	sort.Ints(numbers)

	if len(numbers)%2 == 0 {
		return (float64(numbers[len(numbers)/2]) + float64(numbers[len(numbers)/2-1])) / 2
	} else {
		return float64(numbers[len(numbers)/2])
	}
}

func getMode(numbers []int) int {
	frequency := make(map[int]int)

	for _, number := range numbers {
		frequency[number]++
	}

	mode := numbers[0]
	maxFrequency := frequency[numbers[0]]

	for number, freq := range frequency {
		if freq > maxFrequency {
			mode = number
			maxFrequency = freq
		}
	}

	return mode
}

func getSD(numbers []int) float64 {
	mean := getMean(numbers)

	var sumSquaredDeviations float64

	for _, number := range numbers {
		sumSquaredDeviations += math.Pow(float64(number)-mean, 2)
	}

	return math.Sqrt(sumSquaredDeviations / float64(len(numbers)))
}
