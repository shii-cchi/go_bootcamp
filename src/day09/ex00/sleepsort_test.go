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
		{input: []int{1, 4, 5, 2, 3, 10, 15, 14, 7, 8, 9, 11, 12, 22, 17, 18, 20, 6, 13, 16, 19, 21}, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}},
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
