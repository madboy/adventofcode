package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type key struct {
	up, down, left, right string
}

var keypad = make(map[string]key)

func move(current string, direction string) string {
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

func main() {
	input, err := os.Open("input/2")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	keypad["1"] = key{up: "1", down: "3", left: "1", right: "1"}
	keypad["2"] = key{up: "2", down: "6", left: "2", right: "3"}
	keypad["3"] = key{up: "1", down: "7", left: "2", right: "4"}
	keypad["4"] = key{up: "4", down: "8", left: "3", right: "4"}
	keypad["5"] = key{up: "5", down: "5", left: "5", right: "6"}
	keypad["6"] = key{up: "2", down: "A", left: "5", right: "7"}
	keypad["7"] = key{up: "3", down: "B", left: "6", right: "8"}
	keypad["8"] = key{up: "4", down: "7", left: "C", right: "9"}
	keypad["9"] = key{up: "9", down: "9", left: "8", right: "9"}
	keypad["A"] = key{up: "6", down: "A", left: "A", right: "B"}
	keypad["B"] = key{up: "7", down: "D", left: "A", right: "C"}
	keypad["C"] = key{up: "8", down: "C", left: "B", right: "C"}
	keypad["D"] = key{up: "B", down: "D", left: "D", right: "D"}

	pos := "5"
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			pos = move(pos, string(c))
		}
		fmt.Printf("key: %s\n", pos)
	}
}
