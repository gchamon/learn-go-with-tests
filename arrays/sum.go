package arrays

func Sum(numbers []int) int {
	result := 0

	for _, i := range numbers {
		result += i
	}

	return result
}

func SumAll(arrayOfSlices ...[]int) []int {
	return nil
}
