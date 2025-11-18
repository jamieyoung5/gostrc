package sliceutil

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{4, 3, 2, 1}
	result := Reverse(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}

	if input[0] != 1 {
		t.Error("Original slice was modified")
	}
}

func TestEqual(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}
	d := []int{1, 2}

	if !Equal(a, b) {
		t.Error("Equal slices returned false")
	}
	if Equal(a, c) {
		t.Error("Different content returned true")
	}
	if Equal(a, d) {
		t.Error("Different lengths returned true")
	}
}

func TestCountDuplicates(t *testing.T) {
	input := []string{"a", "b", "a", "c", "b", "b"}

	expected := 2
	result := CountDuplicates(input)

	if result != expected {
		t.Errorf("Expected %d duplicates, got %d", expected, result)
	}
}

func TestCounts(t *testing.T) {
	input := []int{1, 1, 2, 3, 3, 3}
	expected := map[int]int{1: 2, 2: 1, 3: 3}
	result := Counts(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestRandomSubset(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}

	inputCopy := make([]int, len(input))
	copy(inputCopy, input)

	subsetSize := 3
	subset := RandomSubset(inputCopy, subsetSize)

	if len(subset) != subsetSize {
		t.Errorf("Expected subset size %d, got %d", subsetSize, len(subset))
	}

	originalMap := Counts(input)
	for _, v := range subset {
		if originalMap[v] == 0 {
			t.Errorf("Subset contains element %d not in original", v)
		}
	}
}
