package sum

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	// lengthOfNumbers := len(numbersToSum)
	// sums = make([]int, lengthOfNumbers)
	for _, numbers := range numbersToSum {
		// sums[i] = Sum(numbers)
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
