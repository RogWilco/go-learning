package arrays

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(collections ...[]int) (sums []int) {
	for _, numbers := range collections {
		sums = append(sums, Sum(numbers))
	}

	return
}
