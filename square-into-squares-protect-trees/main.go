package main

import "fmt"

func Decompose(n int64) []int64 {
	res := []int64{}
	var sqr int64 = 0
	res = append(res, n)
	for len(res) > 0 {
		top := len(res) - 1
		current := res[top]
		res[top] = -1
		res = res[:top]
		sqr += current * current
		for i := current - 1; i > 0; i-- {
			if sqr-i*i >= 0 {
				sqr -= i * i
				res = append(res, i)
				if sqr == 0 {
					order := []int64{}
					for j := len(res) - 1; j >= 0; j-- {
						order = append(order, res[j])
					}
					return order
				}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(Decompose(50))
	fmt.Println(Decompose(5))
	fmt.Println(Decompose(2))
	fmt.Println(Decompose(7))
}
