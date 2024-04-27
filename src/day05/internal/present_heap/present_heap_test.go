package present_heap

import (
	"errors"
	"sort"
	"testing"
)

var ErrInvalidN = errors.New("invalid value of n")

func TestGetNCoolestPresents_LargeN(t *testing.T) {
	n := 5
	presents := []Present{
		{Value: 5, Size: 1},
		{Value: 4, Size: 5},
		{Value: 3, Size: 1},
		{Value: 5, Size: 2},
	}

	expectedError := ErrInvalidN

	result, err := getNCoolestPresents(n, presents)

	checkResultAndError(t, nil, expectedError, result, err)
}

func TestGetNCoolestPresents_NegativeN(t *testing.T) {
	n := -1
	presents := []Present{
		{Value: 5, Size: 1},
		{Value: 4, Size: 5},
		{Value: 3, Size: 1},
		{Value: 5, Size: 2},
	}

	expectedError := ErrInvalidN

	result, err := getNCoolestPresents(n, presents)

	checkResultAndError(t, nil, expectedError, result, err)
}

func TestGetNCoolestPresents_SmallN1(t *testing.T) {
	n := 2
	presents := []Present{
		{Value: 5, Size: 1},
		{Value: 4, Size: 5},
		{Value: 3, Size: 1},
		{Value: 5, Size: 2},
	}
	expectedResult := []Present{
		{Value: 5, Size: 1},
		{Value: 5, Size: 2},
	}

	result, err := getNCoolestPresents(n, presents)

	checkResultAndError(t, expectedResult, nil, result, err)
}

func TestGetNCoolestPresents_SmallN2(t *testing.T) {
	n := 3
	presents := []Present{
		{Value: 5, Size: 1},
		{Value: 4, Size: 5},
		{Value: 3, Size: 2},
		{Value: 1, Size: 1},
		{Value: 2, Size: 7},
		{Value: 7, Size: 2},
		{Value: 5, Size: 3},
	}
	expectedResult := []Present{
		{Value: 7, Size: 2},
		{Value: 5, Size: 3},
		{Value: 5, Size: 1},
	}

	result, err := getNCoolestPresents(n, presents)

	checkResultAndError(t, expectedResult, nil, result, err)
}

func TestGetNCoolestPresents_EqualN1(t *testing.T) {
	n := 4
	presents := []Present{
		{Value: 5, Size: 1},
		{Value: 4, Size: 5},
		{Value: 3, Size: 1},
		{Value: 5, Size: 2},
	}
	expectedResult := []Present{
		{Value: 5, Size: 1},
		{Value: 5, Size: 2},
		{Value: 4, Size: 5},
		{Value: 3, Size: 1},
	}

	result, err := getNCoolestPresents(n, presents)

	checkResultAndError(t, expectedResult, nil, result, err)
}

func TestGetNCoolestPresents_EqualN2(t *testing.T) {
	n := 7
	presents := []Present{
		{Value: 5, Size: 1},
		{Value: 4, Size: 5},
		{Value: 4, Size: 1},
		{Value: 2, Size: 1},
		{Value: 2, Size: 2},
		{Value: 7, Size: 2},
		{Value: 5, Size: 3},
	}
	expectedResult := []Present{
		{Value: 7, Size: 2},
		{Value: 5, Size: 3},
		{Value: 5, Size: 1},
		{Value: 4, Size: 5},
		{Value: 4, Size: 1},
		{Value: 2, Size: 2},
		{Value: 2, Size: 1},
	}

	result, err := getNCoolestPresents(n, presents)

	checkResultAndError(t, expectedResult, nil, result, err)
}

func checkResultAndError(t *testing.T, expectedResult []Present, expectedError error, result []Present, err error) {
	if err != expectedError {
		if err.Error() != expectedError.Error() {
			t.Errorf("Expected error: %v, got: %v", expectedError, err)
		}
	}

	if expectedError == nil {
		sort.Slice(expectedResult, func(i, j int) bool {
			if expectedResult[i].Value != expectedResult[j].Value {
				return expectedResult[i].Value > expectedResult[j].Value
			}
			return expectedResult[i].Size < expectedResult[j].Size
		})

		if len(result) != len(expectedResult) {
			t.Errorf("Expected length: %d, got: %d", len(expectedResult), len(result))
		}

		for i := range result {
			if result[i] != expectedResult[i] {
				t.Errorf("Expected: %v, got: %v", expectedResult[i], result[i])
			}
		}
	}
}
