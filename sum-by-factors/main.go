package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func PrimeFactors(n int) (pfs []int) {
	if n < 0 {
		n *= -1
	}
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}

func SumOfDivided(lst []int) string {
	result := make(map[int]int)
	for _, v := range lst {
		repeat := 0
		pfs := PrimeFactors(v)
		for _, f := range pfs {
			if repeat == 0 || math.Abs(float64(repeat))-math.Abs(float64(f)) != 0 {
				result[f] += v
			}
			repeat = f
		}
	}

	var keys []int
	for k, _ := range result {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	s := ""
	for _, v := range keys {
		s += "(" + strconv.Itoa(v) + " " + strconv.Itoa(result[v]) + ")"
	}
	return s
}

func main() {
	//
	lst1 := []int{12, -15}
	fmt.Println(SumOfDivided(lst1))
	lst2 := []int{15, 21, 24, 30, 45}
	fmt.Println(SumOfDivided(lst2))
}
