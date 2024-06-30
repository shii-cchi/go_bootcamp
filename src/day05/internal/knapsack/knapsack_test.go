package knapsack

import (
	"reflect"
	"testing"
)

func TestGrabPresents_OnePresent1(t *testing.T) {
	presents := []Present{{Value: 5, Size: 1}}
	capacity := 4
	expected := []Present{{Value: 5, Size: 1}}

	result := grabPresents(presents, capacity)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("For presents: %v and capacity: %d, expected: %v but got: %v", presents, capacity, expected, result)
	}
}

func TestGrabPresents_OnePresent2(t *testing.T) {
	presents := []Present{{Value: 5, Size: 6}}
	capacity := 4
	var expected []Present

	result := grabPresents(presents, capacity)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("For presents: %v and capacity: %d, expected: %v but got: %v", presents, capacity, expected, result)
	}
}

func TestGrabPresents_TwoPresents1(t *testing.T) {
	presents := []Present{{Value: 5, Size: 1}, {Value: 3, Size: 3}}
	capacity := 4
	expected := []Present{{Value: 3, Size: 3}, {Value: 5, Size: 1}}

	result := grabPresents(presents, capacity)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("For presents: %v and capacity: %d, expected: %v but got: %v", presents, capacity, expected, result)
	}
}

func TestGrabPresents_TwoPresents2(t *testing.T) {
	presents := []Present{{Value: 5, Size: 1}, {Value: 3, Size: 3}}
	capacity := 3
	expected := []Present{{Value: 5, Size: 1}}

	result := grabPresents(presents, capacity)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("For presents: %v and capacity: %d, expected: %v but got: %v", presents, capacity, expected, result)
	}
}

func TestGrabPresents_ThreePresents(t *testing.T) {
	presents := []Present{{Value: 5, Size: 1}, {Value: 3, Size: 3}, {Value: 2, Size: 2}}
	capacity := 4
	expected := []Present{{Value: 3, Size: 3}, {Value: 5, Size: 1}}

	result := grabPresents(presents, capacity)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("For presents: %v and capacity: %d, expected: %v but got: %v", presents, capacity, expected, result)
	}
}

func TestGrabPresents_SevenPresents(t *testing.T) {
	presents := []Present{
		{Value: 5, Size: 1},
		{Value: 3, Size: 3},
		{Value: 2, Size: 2},
		{Value: 8, Size: 5},
		{Value: 4, Size: 3},
		{Value: 6, Size: 4},
		{Value: 7, Size: 2},
	}
	capacity := 10
	expected := []Present{
		{Value: 7, Size: 2},
		{Value: 8, Size: 5},
		{Value: 2, Size: 2},
		{Value: 5, Size: 1},
	}

	result := grabPresents(presents, capacity)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("For presents: %v and capacity: %d, expected: %v but got: %v", presents, capacity, expected, result)
	}
}
