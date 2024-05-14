package ex00

import (
	"errors"
	"testing"
)

func TestGetElement(t *testing.T) {
	tests := []struct {
		arr      []int
		idx      int
		expected int
		err      error
	}{
		{[]int{1, 2, 3, 4, 5}, 0, 1, nil},
		{[]int{1, 2, 3, 4, 5}, 2, 3, nil},
		{[]int{1, 2, 3, 4, 5}, 4, 5, nil},
		{[]int{1, 2, 3, 4, 5}, 5, 0, errors.New("error: index is out of bounds")},
		{[]int{1, 2, 3, 4, 5}, -1, 0, errors.New("error: negative index")},
		{[]int{}, 0, 0, errors.New("error: empty slice")},
	}

	for _, test := range tests {
		value, err := getElement(test.arr, test.idx)
		if (err != nil && err.Error() != test.err.Error()) || (err == nil && test.err != nil) {
			t.Errorf("For arr: %v and idx: %d; expected error %v, but got %v", test.arr, test.idx, test.err, err)
		}
		if value != test.expected {
			t.Errorf("For arr: %v and idx: %d; expected value %d, but got %d", test.arr, test.idx, test.expected, value)
		}
	}
}
