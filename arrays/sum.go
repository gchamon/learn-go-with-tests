package arrays

import (
	"github.com/koss-null/funcfrog/pkg/pipe"
)

func Sum(numbers []int) int {
	result := pipe.Slice(numbers).Reduce(func(x, y *int) int { return *x + *y })

	return *result
}

func SumAll(arrayOfSlices ...[]int) []int {
	var sumOfSlices []int

	for _, slice := range arrayOfSlices {
		sumOfSlices = append(sumOfSlices, Sum(slice))
	}

	return sumOfSlices
}
