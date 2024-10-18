package arrays

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(collections ...[]int) (sums []int) {
	sums = make([]int, len(collections))

	for i, numbers := range collections {
		sums[i] = Sum(numbers)
	}

	return
}
