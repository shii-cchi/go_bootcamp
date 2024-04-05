package tests

import (
	"day00/internal/statistics"
	"math"
	"testing"
)

func TestGetMean(t *testing.T) {
	tests := []struct {
		numbers       []int
		expectedMean  float64
		testCaseLabel string
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0, "Test 1: Mean of {1, 2, 3, 4, 5}"},
		{[]int{1, 2}, 1.5, "Test 2: Mean of {1, 2}"},
		{[]int{1}, 1.0, "Test 3: Mean of {1}"},
	}

	for _, tc := range tests {
		actualMean := statistics.GetMean(tc.numbers)
		if actualMean != tc.expectedMean {
			t.Errorf("%s: Expected mean: %v, got: %v", tc.testCaseLabel, tc.expectedMean, actualMean)
		}
	}
}

func TestGetMedian(t *testing.T) {
	tests := []struct {
		numbers        []int
		expectedMedian float64
		testCaseLabel  string
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0, "Test 1: Median of {1, 2, 3, 4, 5}"},
		{[]int{1, 2, 3, 4, 5, 6}, 3.5, "Test 2: Median of {1, 2, 3, 4, 5, 6}"},
		{[]int{5, 3, 1, 4, 2}, 3.0, "Test 3: Median of {5, 3, 1, 4, 2}"},
		{[]int{5, 3, 1, 4, 2, 6}, 3.5, "Test 4: Median of {5, 3, 1, 4, 2, 6}"},
		{[]int{1, 2}, 1.5, "Test 5: Median of {1, 2}"},
		{[]int{1}, 1.0, "Test 6: Median of {1}"},
	}

	for _, tc := range tests {
		actualMedian := statistics.GetMean(tc.numbers)
		if actualMedian != tc.expectedMedian {
			t.Errorf("%s: Expected median: %v, got: %v", tc.testCaseLabel, tc.expectedMedian, actualMedian)
		}
	}
}

func TestGetMode(t *testing.T) {
	tests := []struct {
		numbers       []int
		expectedMode  int
		testCaseLabel string
	}{
		{[]int{1, 2, 3, 4, 5}, 1, "Test 1: Mode of {1, 2, 3, 4, 5}"},
		{[]int{1, 1, 1, 2, 2, 2}, 1, "Test 2: Mode of {1, 1, 1, 2, 2, 2}"},
		{[]int{1}, 1, "Test 3: Mode of {1}"},
	}

	for _, tc := range tests {
		actualMode := statistics.GetMode(tc.numbers)
		if actualMode != tc.expectedMode {
			t.Errorf("%s: Expected median: %v, got: %v", tc.testCaseLabel, tc.expectedMode, actualMode)
		}
	}
}

func TestGetSD(t *testing.T) {
	tests := []struct {
		numbers       []int
		expectedSD    float64
		testCaseLabel string
	}{
		{[]int{1, 2, 3, 4, 5}, 1.41, "Test 1: SD of {1, 2, 3, 4, 5}"},
		{[]int{1, 1, 1, 2, 2, 2}, 0.5, "Test 2: SD of {1, 1, 1, 2, 2, 2}"},
		{[]int{1}, 0, "Test 3: SD of {1}"},
	}

	for _, tc := range tests {
		actualSD := statistics.GetSD(tc.numbers)
		if !floatsAreEqual(actualSD, tc.expectedSD) {
			t.Errorf("%s: Expected SD: %v, got: %v", tc.testCaseLabel, tc.expectedSD, actualSD)
		}
	}
}

func floatsAreEqual(a, b float64) bool {
	const epsilon = 0.01
	return math.Abs(a-b) < epsilon
}
