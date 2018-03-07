package days

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func calculateCoord(n, e int) string {
	tn := strconv.Itoa(n)
	te := strconv.Itoa(e)
	return strings.Join([]string{tn, te}, ",")
}

func calculateMovement(input rune, n, e int) (int, int) {
	switch direction := string(input); direction {
	case "^":
		n++
	case "v":
		n--
	case ">":
		e++
	case "<":
		e--
	}
	return n, e
}

// Run3 for the third day of christmas
func Run3(scanner *bufio.Scanner) string {
	var input string
	for scanner.Scan() {
		input = scanner.Text()
	}
	return fmt.Sprintf("%s\n%s", onlySanta(input), santaAndRoboSanta(input))
}

func onlySanta(input string) string {
	var travels map[string]int
	travels = make(map[string]int)
	n, e := 0, 0
	currentCoord := calculateCoord(n, e)
	travels[currentCoord]++
	for _, c := range input {
		switch direction := string(c); direction {
		case "^":
			n++
		case "v":
			n--
		case ">":
			e++
		case "<":
			e--
		}
		currentCoord = calculateCoord(n, e)
		travels[currentCoord]++
	}
	return fmt.Sprintf("Santa has visited %d houses", len(travels))
}

func santaAndRoboSanta(input string) string {
	var travels map[string]int
	travels = make(map[string]int)
	sn, se, rn, re := 0, 0, 0, 0

	travels[calculateCoord(sn, se)] += 2
	for i, c := range input {
		if (i % 2) == 0 {
			sn, se = calculateMovement(c, sn, se)
			travels[calculateCoord(sn, se)]++
		} else {
			rn, re = calculateMovement(c, rn, re)
			travels[calculateCoord(rn, re)]++
		}
	}
	return fmt.Sprintf("Santa and Robo-Santa have visted %d houses", len(travels))
}
