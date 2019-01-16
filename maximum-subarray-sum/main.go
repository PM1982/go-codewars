package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func MaximumSubarraySum(numbers []int) int {
	res, sum := 0, 0
	for i := range numbers {
		sum += numbers[i]
		res = max(res, sum)
		sum = max(sum, 0)
	}
	return res
}

func main() {
	empty := []int{}
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	b := []int{-2, -1, -3, -4, -1, -2, -1, -5, -4}
	fmt.Println(MaximumSubarraySum(empty))
	fmt.Println(MaximumSubarraySum(a))
	fmt.Println(MaximumSubarraySum(b))
}
