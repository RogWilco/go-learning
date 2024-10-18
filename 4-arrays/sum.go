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

func SumAllTails(collections ...[]int) (sums []int) {
	for _, numbers := range collections {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}
