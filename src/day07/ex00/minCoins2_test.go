package ex00

import (
	"reflect"
	"testing"
)

func TestMinCoins2(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		val      int
		expected []int
	}{
		{
			name:     "Test 1",
			coins:    []int{1, 5, 10},
			val:      13,
			expected: []int{10, 1, 1, 1},
		},
		{
			name:     "Test 2",
			coins:    []int{1, 5, 13},
			val:      13,
			expected: []int{13},
		},
		{
			name:     "Test 3",
			coins:    []int{1, 5},
			val:      13,
			expected: []int{5, 5, 1, 1, 1},
		},
		{
			name:     "Test 4",
			coins:    []int{1},
			val:      6,
			expected: []int{1, 1, 1, 1, 1, 1},
		},
		{
			name:     "Test 5",
			coins:    []int{},
			val:      13,
			expected: []int{},
		},
		{
			name:     "Test 6",
			coins:    []int{},
			val:      0,
			expected: []int{},
		},
		{
			name:     "Test 7",
			coins:    []int{1, 5, 10},
			val:      0,
			expected: []int{},
		},
		{
			name:     "Test 8",
			coins:    []int{1, 5, 10},
			val:      -1,
			expected: []int{},
		},
		{
			name:     "Test 9",
			coins:    []int{10, 5, 1},
			val:      11,
			expected: []int{10, 1},
		},
		{
			name:     "Test 10",
			coins:    []int{10, 5, 1, 2},
			val:      12,
			expected: []int{10, 2},
		},
		{
			name:     "Test 11",
			coins:    []int{1, 5, 10, 10},
			val:      12,
			expected: []int{10, 1, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MinCoins2(test.val, test.coins)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Test %s: Expected %v, but got %v", test.name, test.expected, result)
			}
		})
	}
}
