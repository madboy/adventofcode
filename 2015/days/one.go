package days

import (
	"bufio"
	"fmt"
)

// Run1 return answer for day 1 part 1 2015
func Run1(scanner *bufio.Scanner) string {
	var input string
	for scanner.Scan() {
		input = scanner.Text()
	}
	floor := 0
	for _, c := range input {
		switch char := string(c); char {
		case "(":
			floor++
		case ")":
			floor--
		}
	}
	return fmt.Sprintf("Santa ends up on floor: %d,\nand basement entrance position is: %d", santasFloor(input), basementPosition(input))
}

func santasFloor(input string) int {
	floor := 0
	for _, c := range input {
		switch char := string(c); char {
		case "(":
			floor++
		case ")":
			floor--
		}
	}
	return floor
}

func basementPosition(input string) int {
	floor := 0
	pos := 0 //reason for using pos is if there's any non () in the input
	for _, c := range input {
		if string(c) == "(" {
			floor++
			pos++
		} else if string(c) == ")" {
			floor--
			pos++
		}
		if floor == -1 {
			return pos
		}
	}
	return 0
}
