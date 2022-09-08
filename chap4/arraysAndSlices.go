package arraysAndSlices

func Sum(numbers []int) (sum int) {
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return
}

func SumAllTails(numberToSum ...[]int) (sums []int) {
	sums = make([]int, len(numberToSum))
	for i, numbers := range numberToSum {
		tail := numbers[1:]
		sums[i] = Sum(tail)
	}
	return
}
