package arrays

func Sum(numbers [5]int) int {
	result := 0

	for _, i := range numbers {
		result += i
	}

	return result
}
