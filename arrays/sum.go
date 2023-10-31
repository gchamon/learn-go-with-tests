package arrays

import (
	"github.com/koss-null/funcfrog/pkg/pipe"
)

func Sum(numbers []int) int {
	result := pipe.Slice(numbers).Reduce(func(x, y *int) int { return *x + *y })

	return *result
}

func SumAll(arrayOfSlices ...[]int) []int {
	result := pipe.
		Map(pipe.Slice(arrayOfSlices), Sum).
		Do()

	return result
}
