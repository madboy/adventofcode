package days

import (
	"bufio"
	"fmt"
)

// Formula for calculating the next code on the paper
func nextCode(current int) int {
	return (current * 252533) % 33554393
}

func calculateTargetValue(row, col int) int {
	// start values per row
	rv := 1
	for r := 2; r <= row; r++ {
		rv = r - 1 + rv
	}

	// column progression for row
	cv := rv
	r := row
	for c := 2; c <= col; c++ {
		cv = (c - 1) + r + cv

	}
	return cv
}

// Run25 in which we are fighting the boss with magic
func Run25(scanner *bufio.Scanner) string {
	targetRow := 2981
	targetCol := 3075
	targetV := calculateTargetValue(targetRow, targetCol)

	column := 1
	row := 1
	start := 20151125
	curCode := start
	for c := 0; c < targetV-1; c++ {
		if row == 1 {
			// move down to next diagonal
			row = column + 1
			column = 1
		} else {
			// move diagonally up right
			row--
			column++
		}
		curCode = nextCode(curCode)

	}

	return fmt.Sprintf("The unlock code for the weather machine is: %d", curCode)
}
