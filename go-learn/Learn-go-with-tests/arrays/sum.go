package arrays

func Sum(array []int) int {
	sum := 0
	for _, i := range array {
		sum += i
	}
	return sum
}

func SumAll(numebrsToSum ...[]int) []int {
	//lengthOfNmbers := len(numebrsToSum)
	//sums = make([]int, lengthOfNmbers)
	var sums []int
	for _, numbers := range numebrsToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers[1:]))
		}
	}
	return sums

}
