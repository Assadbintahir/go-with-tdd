package slices_arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("should return the sum", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		expected := 6

		if got != expected {
			t.Errorf("got %d but expected %d given %v", got, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	checkSum := func(t *testing.T, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %d but expected %d", got, expected)
		}
	}

	t.Run("should return the sum slice when multiple slices are provided", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{4, 5})
		expected := []int{3, 9}
		checkSum(t, got, expected)
	})

	t.Run("should return the sum of tails only", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{4, 5})
		expected := []int{2, 5}
		checkSum(t, got, expected)
	})

	t.Run("should safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}
		checkSum(t, got, expected)
	})
}
