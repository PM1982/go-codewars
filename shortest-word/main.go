package main

import (
	"fmt"
	"math"
)

func FindShort(s string) int {
	// your code
	minLen, curLen := math.MaxFloat64, .0

	for i := 0; i < len(s); i++ {
		if s[i:i+1] == " " {
			minLen = math.Min(minLen, curLen)
			curLen = .0
			continue
		}
		curLen++
	}
	minLen = math.Min(minLen, curLen)
	return int(minLen)
}

func main() {
	a := "bitcoin take over the world maybe who knows perhaps"
	b := "turns out random test cases are easier than writing out basic ones"
	fmt.Println(FindShort(a))
	fmt.Println(FindShort(b))
}
