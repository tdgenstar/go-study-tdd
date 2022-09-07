package arraysAndSlices

func Sum(numbers []int) (sum int) {
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return
}

func SumAll(numbersToSum ...[]int) []int {
	return []int{}
}
