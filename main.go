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
	}
	return -1, -1
}

func isNumValid(puzz [][]int, guess int, row int, column int) bool {
	for index := range puzz {
		if puzz[row][index] == guess && column != index {
			return false
		}
	}

	for index := range puzz {
		if puzz[index][column] == guess && row != index {
			return false
		}
	}

	// Is number valid for box
	// row 0 & column 3
	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			if puzz[row-row%3+k][column-column%3+k] == guess && (row-row%3+k != row || column-column%3+l != column) {
				return false
			}
		}
	}
	return true
}

func solvePuzzle(puzz [][]int) bool {
	row, column := getEmptySpace(puzz)
	if row == -1 {
		return true
	} else {
		row, column = getEmptySpace(puzz)
	}
	for i := 1; i <= 9; i++ {
		if isNumValid(puzz, i, row, column) {
			puzz[row][column] = i

			if solvePuzzle(puzz) {
				return true
			}
			puzz[row][column] = 0
		}
	}
	return false
}

func main() {
	puzz := [][]int{
		{0, 0, 0, 0, 3, 5, 0, 7, 0},
		{2, 5, 0, 0, 4, 6, 8, 0, 1},
		{0, 1, 3, 7, 0, 8, 0, 4, 9},
		{1, 9, 0, 0, 0, 7, 0, 0, 4},
		{0, 0, 5, 0, 0, 2, 0, 9, 6},
		{8, 0, 2, 0, 9, 4, 0, 0, 7},
		{3, 7, 0, 0, 0, 9, 0, 0, 0},
		{0, 6, 1, 0, 7, 0, 0, 0, 0},
		{4, 0, 0, 5, 8, 1, 0, 0, 0},
	}

	displayBoard(puzz)

	row, _ := getEmptySpace(puzz)
	if row != -1 {
		fmt.Println(getEmptySpace(puzz))
	} else {
		fmt.Println("Puzzle is solved")
	}

	fmt.Println(isNumValid(puzz, 7, 4, 0))

	displayBoard(puzz)
	fmt.Println()
	solvePuzzle(puzz)
	displayBoard(puzz)
}
