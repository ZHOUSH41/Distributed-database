package arrays

func Sum(array []int) int {
	sum := 0
	for _,i:= range(array) {
		sum += i
	}
	return sum
}

func SumAll(numebrsToSum ...[]int) (sums []int){
	lengthOfNmbers := len(numebrsToSum)
	sums = make([]int, lengthOfNmbers)
	for i, numbers := range numebrsToSum {
		sums[i] = Sum(numbers)
	}
	return
}
