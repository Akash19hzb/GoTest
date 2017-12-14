package problems

import "testing"

func TestShouldReturnTrueWhenThreeNumbersAreEqual(t *testing.T) {
	areThreeNumbersEqual := equalOrNotEqual(1, 1, 1)
	if !areThreeNumbersEqual {
		t.Error("Expected true, got", areThreeNumbersEqual)
	}
}

func TestShouldReturnFalseWhenThreeNumbersAreNotEqual(t *testing.T) {

	areThreeNumbersEqual := equalOrNotEqual(1, 2, 3)
	if areThreeNumbersEqual {
		t.Error("Expected false, got", areThreeNumbersEqual)
	}
}