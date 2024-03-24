package logic

import (
	"testing"
)

// Test for getRemainingItems function
func TestGetRemainingItems(t *testing.T) {
	items := []string{"item1", "item2", "item3"}
	value := "item2"

	result := getRemainingItems(value, items)
	expected := []string{"item1", "item3"}

	if !sliceEqual(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

// Helper function to check if two slices are equal
func sliceEqual(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
