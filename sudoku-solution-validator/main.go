package main

import "fmt"

var testTrue = [][]int{
	{5, 3, 4, 6, 7, 8, 9, 1, 2},
	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	{1, 9, 8, 3, 4, 2, 5, 6, 7},
	{8, 5, 9, 7, 6, 1, 4, 2, 3},
	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	{9, 6, 1, 5, 3, 7, 2, 8, 4},
	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	{3, 4, 5, 2, 8, 6, 1, 7, 9},
}

var testFalse = [][]int{
	{5, 3, 4, 6, 7, 8, 9, 1, 2},
	{6, 7, 2, 1, 9, 0, 3, 4, 8},
	{1, 0, 0, 3, 4, 2, 5, 6, 0},
	{8, 5, 9, 7, 6, 1, 0, 2, 0},
	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	{9, 0, 1, 5, 3, 7, 2, 1, 4},
	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	{3, 0, 0, 4, 8, 1, 1, 7, 9},
}

func ValidateSolution(sudoku [][]int) bool {
	for r := 0; r < 9; r++ {
		flagrow := 0
		flagcol := 0
		flagbox := 0
		for c := 0; c < 9; c++ {
			if sudoku[r][c] < 1 || sudoku[r][c] > 9 || sudoku[c][r] < 1 || sudoku[c][r] > 9 {
				return false
			}

			if flagrow&(1<<uint(sudoku[r][c])) != 0 {
				return false
			}

			if flagcol&(1<<uint(sudoku[c][r])) != 0 {
				return false
			}

			i := (r%3)*3 + c%3
			j := (r/3)*3 + c/3
			if flagbox&(1<<uint(sudoku[j][i])) != 0 {
				return false
			}
			flagrow |= 1 << uint(sudoku[r][c])
			flagcol |= 1 << uint(sudoku[c][r])
			flagbox |= 1 << uint(sudoku[j][i])
		}
	}
	return true
}

func main() {
	fmt.Println(ValidateSolution(testTrue))
	fmt.Println(ValidateSolution(testFalse))
}
