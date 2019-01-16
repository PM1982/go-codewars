package main

import "fmt"

func calc(n int) ([]int, []int) {
	ann := make([]int, n, n)
	john := make([]int, n, n)
	ann[0] = 1
	john[0] = 0
	for i := 1; i < n; i++ {
		ann[i] = i - john[ann[i-1]]
		john[i] = i - ann[john[i-1]]
	}
	return ann, john
}

func Ann(n int) []int {
	res, _ := calc(n)
	return res
}
func John(n int) []int {
	_, res := calc(n)
	return res
}
func SumJohn(n int) int {
	_, res := calc(n)
	sum := 0
	for i := 0; i < n; i++ {
		sum += res[i]
	}
	return sum
}
func SumAnn(n int) int {
	res, _ := calc(n)
	sum := 0
	for i := 0; i < n; i++ {
		sum += res[i]
	}
	return sum
}

func main() {
	fmt.Println(John(11))
	fmt.Println(Ann(6))
	fmt.Println(SumJohn(75))
	fmt.Println(SumAnn(115))
}
