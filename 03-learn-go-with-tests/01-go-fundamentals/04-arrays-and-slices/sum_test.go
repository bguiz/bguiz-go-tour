package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestBothSumFunctions(t *testing.T) {

	t.Run("fixed size", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum5Array(numbers)
		want := 15

		if want != got {
			t.Errorf("want %d got %d given %v", want, got, numbers)
		}
	})

	t.Run("dynamic size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := SumSlice(numbers)
		want := 15

		if want != got {
			t.Errorf("want %d got %d given %v", want, got, numbers)
		}
	})

}

// NOTE
// - ran code coverage tool using `go test -cover`

func TestSumAll(t *testing.T) {

	t.Run("SumAll 2x2", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 4})
		want := []int{3, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %d got %d", want, got)
		}
	})

	t.Run("SumAllUsingAppend 2x2", func(t *testing.T) {
		got := SumAllUsingAppend([]int{1, 2}, []int{3, 4})
		want := []int{3, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %d got %d", want, got)
		}
	})

}

// NOTE
// - use `reflect.DeepEqual` to compare slices, because the equality operator (`==`)
//   only is meant for primitive types
//   - also, this function is not type safe

func TestSumAllTails(t *testing.T) {

	t.Run("SumAllTails 2x3", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 1, 2})
		want := []int{5, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %d got %d", want, got)
		}
	})

	t.Run("SumAllTails includes empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{9, 1, 2})
		want := []int{0, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %d got %d", want, got)
		}
	})

}
