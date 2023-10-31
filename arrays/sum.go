package arrays

func Sum(numbers []int) int {
	result := 0

	for _, i := range numbers {
		result += i
	}

	return result
}

func SumAll(arrayOfSlices ...[]int) []int {
	var sumOfSlices []int

	for _, slice := range arrayOfSlices {
		sumOfSlices = append(sumOfSlices, Sum(slice))
	}

	return sumOfSlices
}
