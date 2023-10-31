package arrays

import "github.com/mariomac/gostream/stream"

func Sum(numbers []int) int {
	result, _ := stream.OfSlice(numbers).Reduce(func(a, b int) int { return a + b })

	return result
}

func SumAll(arrayOfSlices ...[]int) []int {
	var sumOfSlices []int

	for _, slice := range arrayOfSlices {
		sumOfSlices = append(sumOfSlices, Sum(slice))
	}

	return sumOfSlices
}
