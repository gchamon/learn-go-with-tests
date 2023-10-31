package arrays

import (
	"learn/util"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("array of five numbers", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		util.AssertCorrectMessageIntGivenArray(t, got, want, numbers)
	})
}
