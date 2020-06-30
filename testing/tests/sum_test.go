package tests

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 4, 5, 3}
	expected := 15
	actual := Sum(numbers...)

	if actual != expected {
		t.Errorf("Expected the sum of %v to be %d but instead got %d", numbers, expected, actual)
	}
}
