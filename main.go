package main

import (
	"fmt"
	"strconv"
)

// Output the current board to screen
func displayBoard(puzz [][]int) {
	puzz_cols := len(puzz[0])
	//cycles the row in the puzzle
	for i := 0; i < len(puzz); i++ {
		//cycles the column in the puzzle
		for j := 0; j < puzz_cols; j++ {
			fmt.Print(strconv.Itoa(puzz[i][j]) + " ")
		}
		fmt.Println()
	}
}

// Return a valid row and column for empty space
func getEmptySpace(puzz [][]int) (int, int) {
	puzz_cols := len(puzz[0])
	//cycles the row in the puzzle
	for i := 0; i < len(puzz); i++ {
		//cycles the column in the puzzle
		for j := 0; j < puzz_cols; j++ {
			if puzz[i][j] == 0 {
				return i, j
			}
		}
		fmt.Println()
	}
	return -1, -1
}

func main() {
	puzz := [][]int{
		{1, 0, 0, 0, 3, 5, 0, 7, 0},
		{2, 5, 0, 0, 4, 6, 8, 0, 1},
		{0, 1, 3, 7, 0, 8, 0, 4, 9},
		{1, 9, 0, 0, 0, 7, 0, 0, 4},
		{0, 0, 5, 0, 0, 2, 0, 9, 6},
		{8, 0, 2, 0, 9, 4, 0, 0, 7},
		{3, 7, 0, 0, 0, 9, 0, 0, 0},
		{0, 6, 1, 0, 7, 0, 0, 0, 0},
		{4, 0, 0, 5, 8, 1, 0, 0, 1},
	}

	displayBoard(puzz)

	row, _ := getEmptySpace(puzz)
	if row != -1 {
		fmt.Println(getEmptySpace(puzz))
	} else {
		fmt.Println("Puzzle is solved")
	}
}
