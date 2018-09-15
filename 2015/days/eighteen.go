package days

import (
	"bufio"
	"fmt"
)

func getLight(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func getNeighboursOn(x, y int, display map[string]bool) int {
	neighbours := []string{
		getLight(x-1, y-1), getLight(x, y-1), getLight(x+1, y-1),
		getLight(x-1, y), getLight(x+1, y),
		getLight(x-1, y+1), getLight(x, y+1), getLight(x+1, y+1),
	}
	nbrOn := 0
	for _, n := range neighbours {
		if display[n] {
			nbrOn++
		}
	}
	return nbrOn
}

func resetCorners(display map[string]bool, dSize int) {
	for _, c := range []string{getLight(0, 0), getLight(0, dSize-1), getLight(dSize-1, 0), getLight(dSize-1, dSize-1)} {
		display[c] = true
	}
}

func printDisplay(display map[string]bool, dSize int) {
	out := ""
	for y := 0; y < dSize; y++ {
		for x := 0; x < dSize; x++ {
			if display[getLight(x, y)] {
				out += "#"
			} else {
				out += "."
			}
		}
		out += "\n"
	}
	fmt.Println(out)
}

func lightsOn(display map[string]bool) int {
	nbrOn := 0
	for _, v := range display {
		if v {
			nbrOn++
		}
	}
	return nbrOn
}

func updateDisplay(display map[string]bool, dSize int, broken bool) map[string]bool {
	tDisplay := make(map[string]bool)
	for x := 0; x < dSize; x++ {
		for y := 0; y < dSize; y++ {
			nbrOn := getNeighboursOn(x, y, display)
			if display[getLight(x, y)] {
				if nbrOn == 2 || nbrOn == 3 {
					tDisplay[getLight(x, y)] = true
				} else {
					tDisplay[getLight(x, y)] = false
				}
			} else {
				if nbrOn == 3 {
					tDisplay[getLight(x, y)] = true
				}
			}
		}
	}

	if broken {
		resetCorners(tDisplay, dSize)
	}
	return tDisplay
}

// Run18 in which we turn it off and on again
func Run18(scanner *bufio.Scanner) string {
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	dSize := len(input[0])
	display := make(map[string]bool)
	brokenDisplay := make(map[string]bool)
	for y, line := range input {
		for x, c := range line {
			display[getLight(x, y)] = c == 35       // code for #
			brokenDisplay[getLight(x, y)] = c == 35 // code for #
		}
	}

	resetCorners(brokenDisplay, dSize)

	for i := 0; i < 100; i++ {
		display = updateDisplay(display, dSize, false)
		brokenDisplay = updateDisplay(brokenDisplay, dSize, true)
	}

	return fmt.Sprintf("After power cycling for a while %d and %d lights are on", lightsOn(display), lightsOn(brokenDisplay))
}
