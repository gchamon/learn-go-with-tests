package arrays

import (
	"github.com/koss-null/funcfrog/pkg/pipe"
)

func sum(x, y *int) int {
	return *x + *y
}

func fillSlice(slice []int) []int {
	if len(slice) < 2 {
		return append(slice, []int{0, 0}...)
	}
	return slice
}

func Sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	result := pipe.Slice(numbers).Reduce(sum)

	return *result
}

func SumAll(arrayOfSlices ...[]int) []int {
	result := pipe.
		Map(pipe.Slice(arrayOfSlices).Map(fillSlice), Sum).
		Do()

	return result
}

func tail(slice []int) []int {
	return slice[1:]
}

func SumAllTails(arrayOfSlices ...[]int) []int {
	pipeArrayOfTails := pipe.Map(pipe.Slice(arrayOfSlices).Map(fillSlice), tail)

	return pipe.Map(pipeArrayOfTails, Sum).Do()
}

func SumIt(numbers []int) int {
	result := 0

	for _, n := range numbers {
		result += n
	}

	return result
}

func SumAllIt(arrayOfSlices ...[]int) []int {
	var result []int

	for _, numbers := range arrayOfSlices {
		result = append(result, SumIt(fillSlice(numbers)))
	}

	return result
}

func SumAllTailsIt(arrayOfSlices ...[]int) []int {
	var result []int

	for _, numbers := range arrayOfSlices {
		result = append(result, SumIt(tail(fillSlice(numbers))))
	}

	return result
}
