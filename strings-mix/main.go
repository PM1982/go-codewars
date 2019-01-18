package main

import (
	"fmt"
)

type order struct {
	v1  int
	v2  int
	asc int
	idx int
	s   string
	ch  string
}

func (o *order) Max() int {
	if o.v1 > o.v2 {
		return o.v1
	}
	return o.v2
}

func (o *order) CalcIdx() *order {
	isF := 1
	o.s = "2:"
	idx := o.Max()*10000 - o.asc
	diff := o.v1 - o.v2
	if diff > 0 {
		isF = 2
		o.s = "1:"
	}
	if diff == 0 {
		isF = 0
		o.s = "=:"
	}
	idx += isF * 1000
	o.idx = idx
	return o
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

	for k, v := range s1Amount {
		r := []rune(k)
		res[k] = order{v, res[k].v2, int(r[0]), 0, "", k}
	}

	for k, v := range s2Amount {
		r := []rune(k)
		res[k] = order{res[k].v1, v, int(r[0]), 0, "", k}
	}

	for k, v := range res {
		res[k] = *v.CalcIdx()
		if v.v1 < 2 && v.v2 < 2 {
			delete(res, k)
		}
	}

	s := ""
	for len(res) > 0 {
		max := 0
		idx := ""
		for k, v := range res {
			if v.idx > max {
				max = v.idx
				idx = k
			}
		}
		v := res[idx]
		s += v.s
		for i := v.Max(); i > 0; i-- {
			s += v.ch
		}
		if len(res) > 1 {
			s += "/"
		}
		delete(res, idx)
	}
	return s
}

func main() {
	fmt.Println(Mix("Are they here", "yes, they are here"))
	fmt.Println(Mix("uuuuuu", "uuuuuu"))
	fmt.Println(Mix("looping is fun but dangerous", "less dangerous than coding"))
}
