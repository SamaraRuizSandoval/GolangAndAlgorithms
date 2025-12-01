package basic_syntax

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(slices ...[]int) []int {
	sums := make([]int, len(slices))
	for i, slice := range slices {
		sums[i] = Sum(slice)
	}
	return sums
}

func SumAllTails(slices ...[]int) []int {
	sums := make([]int, len(slices))
	for i, slice := range slices {
		if len(slice) == 0 {
			sums[i] = 0
		} else {
			sums[i] = Sum(slice[1:])
		}

	}
	return sums
}
