package ex00

import (
	"slices"
	"testing"
)

func TestSleepSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{5, 3, 2, 1, 4}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{10, 1, 2, 9, 8, 7, 3, 4, 6, 5}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 5, 5, 5, 5}, expected: []int{5, 5, 5, 5, 5}},
	}

	for _, test := range tests {
		resultChan := sleepSort(test.input)
		var result []int
		for num := range resultChan {
			result = append(result, num)
		}

		if slices.Compare(result, test.expected) != 0 {
			t.Errorf("For input %v, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
