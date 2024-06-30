package present_heap

import (
	"errors"
	"reflect"
	"testing"
)

var ErrInvalidN = errors.New("invalid value of n")

func TestGetNCoolestPresentsNormalCases(t *testing.T) {
	tests := []struct {
		presents []Present
		n        int
		expected []Present
	}{
		{
			presents: []Present{
				{Value: 5, Size: 1},
				{Value: 4, Size: 5},
				{Value: 3, Size: 1},
				{Value: 5, Size: 2},
			},
			n: 2,
			expected: []Present{
				{Value: 5, Size: 1},
				{Value: 5, Size: 2},
			},
		},
		{
			presents: []Present{
				{Value: 1, Size: 1},
				{Value: 1, Size: 2},
				{Value: 1, Size: 3},
				{Value: 1, Size: 4},
			},
			n: 3,
			expected: []Present{
				{Value: 1, Size: 1},
				{Value: 1, Size: 2},
				{Value: 1, Size: 3},
			},
		},
		{
			presents: []Present{
				{Value: 2, Size: 2},
				{Value: 3, Size: 3},
				{Value: 1, Size: 1},
				{Value: 5, Size: 5},
			},
			n: 1,
			expected: []Present{
				{Value: 5, Size: 5},
			},
		},
	}

	for _, test := range tests {
		result, err := getNCoolestPresents(test.n, test.presents)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For presents %v and n %d, expected %v, got %v", test.presents, test.n, test.expected, result)
		}
	}
}

func TestGetNCoolestPresentsErrorCases(t *testing.T) {
	tests := []struct {
		presents []Present
		n        int
	}{
		{
			presents: []Present{
				{Value: 5, Size: 1},
				{Value: 4, Size: 5},
			},
			n: 3,
		},
		{
			presents: []Present{
				{Value: 5, Size: 1},
			},
			n: -1,
		},
	}

	for _, test := range tests {
		_, err := getNCoolestPresents(test.n, test.presents)
		if err == nil {
			t.Errorf("Expected error for presents %v and n %d, but got none", test.presents, test.n)
		}
	}
}
