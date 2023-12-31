package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("array of five numbers", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
	t.Run("array of 1 number", func(t *testing.T) {

		numbers := []int{1}

		got := Sum(numbers)
		want := 1

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
	t.Run("empty array", func(t *testing.T) {

		numbers := []int{}

		got := Sum(numbers)
		want := 0

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	given := [][]int{{1, 2, 3}, {4, 5, 6}}
	got := SumAll(given...)
	want := []int{6, 15}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v, given %v", got, want, given)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("make sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 5, 6})
		want := []int{0, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestSumIt(t *testing.T) {
	t.Run("array of five numbers", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}

		got := SumIt(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
	t.Run("array of 1 number", func(t *testing.T) {

		numbers := []int{1}

		got := SumIt(numbers)
		want := 1

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
	t.Run("empty array", func(t *testing.T) {

		numbers := []int{}

		got := SumIt(numbers)
		want := 0

		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAllIt(t *testing.T) {
	got := SumAllIt([]int{1, 2, 3}, []int{4, 5, 6})
	want := []int{6, 15}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumAllTailsIt(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("make sums of some slices", func(t *testing.T) {
		got := SumAllTailsIt([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}
		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTailsIt([]int{}, []int{4, 5, 6})
		want := []int{0, 11}
		checkSums(t, got, want)
	})
}
