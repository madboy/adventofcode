package days

import (
	"bufio"
	"bytes"
	"fmt"
)

type key struct {
	up, down, left, right string
}

func checkEdge(c, n int, edges []int) int {
	for _, e := range edges {
		if c == e {
			return c
		}
	}
	return n
}

func moveSimple(current int, direction string) int {
	urEdges := []int{0, 1, 2}
	drEdges := []int{6, 7, 8}
	lcEdges := []int{0, 3, 6}
	rcEdges := []int{2, 5, 8}
	rows := 3
	switch direction {
	case "U":
		return checkEdge(current, current-rows, urEdges)
	case "D":
		return checkEdge(current, current+rows, drEdges)
	case "L":
		return checkEdge(current, current-1, lcEdges)
	case "R":
		return checkEdge(current, current+1, rcEdges)
	}
	return current
}

func moveAdvanced(current string, direction string, keypad map[string]key) string {
	switch direction {
	case "U":
		return keypad[current].up
	case "D":
		return keypad[current].down
	case "L":
		return keypad[current].left
	case "R":
		return keypad[current].right
	}
	return current
}

func simpleKeyPad(input []string) string {
	keypad := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var code bytes.Buffer
	pos := 4
	for _, line := range input {
		for _, c := range line {
			pos = moveSimple(pos, string(c))
		}
		code.WriteString(fmt.Sprintf("%d", keypad[pos]))
	}
	return fmt.Sprintf("The bathroom code is: %s", code.String())
}

func advancedKeypad(input []string) string {
	var keypad = make(map[string]key)
	var code bytes.Buffer

	keypad["1"] = key{up: "1", down: "3", left: "1", right: "1"}
	keypad["2"] = key{up: "2", down: "6", left: "2", right: "3"}
	keypad["3"] = key{up: "1", down: "7", left: "2", right: "4"}
	keypad["4"] = key{up: "4", down: "8", left: "3", right: "4"}
	keypad["5"] = key{up: "5", down: "5", left: "5", right: "6"}
	keypad["6"] = key{up: "2", down: "A", left: "5", right: "7"}
	keypad["7"] = key{up: "3", down: "B", left: "6", right: "8"}
	keypad["8"] = key{up: "4", down: "C", left: "7", right: "9"}
	keypad["9"] = key{up: "9", down: "9", left: "8", right: "9"}
	keypad["A"] = key{up: "6", down: "A", left: "A", right: "B"}
	keypad["B"] = key{up: "7", down: "D", left: "A", right: "C"}
	keypad["C"] = key{up: "8", down: "C", left: "B", right: "C"}
	keypad["D"] = key{up: "B", down: "D", left: "D", right: "D"}

	pos := "5"
	for _, line := range input {
		for _, c := range line {
			pos = moveAdvanced(pos, string(c), keypad)
		}
		code.WriteString(pos)
	}
	return fmt.Sprintf("The correct bathroom code is: %s", code.String())
}

// Run2 will help us get some relief
func Run2(scanner *bufio.Scanner) string {
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return fmt.Sprintf("%s\n%s", simpleKeyPad(input), advancedKeypad(input))
}
