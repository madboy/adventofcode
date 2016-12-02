package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var keypad = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
var urEdges = []int{0, 1, 2}
var drEdges = []int{6, 7, 8}
var lcEdges = []int{0, 3, 6}
var rcEdges = []int{2, 5, 8}
var rows = 3

func checkEdge(c, n int, edges []int) int {
	for _, e := range edges {
		if c == e {
			return c
		}
	}
	return n
}

func move(current int, direction string) int {
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

func main() {
	input, err := os.Open("input/2")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	pos := 4
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		for _, c := range line {
			pos = move(pos, string(c))
		}
		fmt.Printf("key: %d\n", keypad[pos])
	}
}
