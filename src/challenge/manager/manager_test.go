package manager

import (
	"reflect"
	"testing"
)

// check for the duplicates values in a slice
func TestInts(t *testing.T) {
	exampleSlice := []int{0, 1, 1, 9, 8}
	result := Ints(exampleSlice)

	expectedSlice := []int{0, 1, 9, 8}

	if reflect.DeepEqual(result, expectedSlice) == false {
		t.Errorf("Resuted slice incorrect, got: %d, want: %d.", expectedSlice, exampleSlice)
	}
}
