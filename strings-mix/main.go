package main

import (
	"fmt"
	"sort"
)

type order struct {
	v1   int
	v2   int
	isEQ bool
}

func (o *order) Max() int {
	if o.v1 > o.v2 {
		return o.v1
	}
	return o.v2
}

func Mix(s1, s2 string) string {
	isLowercase := map[string]bool{
		"a": true, "b": true, "c": true, "d": true, "e": true,
		"f": true, "g": true, "h": true, "i": true, "j": true,
		"k": true, "l": true, "m": true, "n": true, "o": true,
		"p": true, "q": true, "r": true, "s": true, "t": true,
		"u": true, "v": true, "w": true, "x": true, "y": true,
		"z": true,
	}

	s1Amount := make(map[string]int)
	s2Amount := make(map[string]int)

	for i := len(s1) - 1; i >= 0; i-- {
		if isLowercase[string(s1[i])] == true {
			s1Amount[string(s1[i])]++
		}
	}
	for i := len(s2) - 1; i >= 0; i-- {
		if isLowercase[string(s2[i])] == true {
			s2Amount[string(s2[i])]++
		}
	}

	res := map[string]order{}

	fmt.Println(s1Amount)
	fmt.Println(s2Amount)

	for k, v := range s1Amount {
		res[k] = order{v, res[k].v2, v == res[k].v2}
	}
	for k, v := range s2Amount {
		res[k] = order{res[k].v1, v, res[k].v1 == v}
	}
	for k, v := range res {
		if v.v1 < 2 && v.v2 < 2 {
			delete(res, k)
		}
	}

	keys := make([]string, 0, len(res))
	for k := range res {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Print(k, res[k])
	}
	fmt.Println("")

	rNEQ := map[string]int{}
	rEQ := map[string]int{}
	for k, v := range res {
		if v.isEQ != true {
			rNEQ[k] = v.Max()
		} else {
			rEQ[k] = v.v1
		}
	}

	fmt.Println(res)
	fmt.Println(rNEQ)
	fmt.Println(rEQ)
	return ""
}

func main() {
	fmt.Println(Mix("Are they here", "yes, they are here"))
	fmt.Println(Mix("uuuuuu", "uuuuuu"))
	fmt.Println(Mix("looping is fun but dangerous", "less dangerous than coding"))
}
