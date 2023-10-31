package arrays

import (
	"github.com/koss-null/funcfrog/pkg/pipe"
)

func sum(x, y *int) int {
	return *x + *y
}

func Sum(numbers []int) int {
	result := pipe.Slice(numbers).Reduce(sum)

	return *result
}

func SumAll(arrayOfSlices ...[]int) []int {
	result := pipe.
		Map(pipe.Slice(arrayOfSlices), Sum).
		Do()

	return result
}

func tail(slice []int) []int {
	return slice[1:]
}

func SumAllTails(arrayOfSlices ...[]int) []int {
	pipeArrayOfTails := pipe.Map(pipe.Slice(arrayOfSlices), tail)

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
		result = append(result, SumIt(numbers))
	}

	return result
}

func SumAllTailsIt(arrayOfSlices ...[]int) []int {
	var result []int

	for _, numbers := range arrayOfSlices {
		result = append(result, SumIt(tail(numbers)))
	}

	return result
}
