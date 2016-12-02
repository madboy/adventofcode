package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

var (
	facing  point
	coord   point
	twice   point
	origo   point
	visited map[point]int
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// walk each block step by sneaky step
func walk(steps, face int, start point, xwards bool) {
	for i := 0; i < steps; i++ {
		if xwards {
			start.x += face * 1
		} else {
			start.y += face * 1
		}
		p := point{x: start.x, y: start.y}
		visited[p]++
		if visited[p] > 1 && (twice == origo) {
			twice = p
		}
	}
}

// newCoord updates the current to where the instrcution tells us to go
func newCoord(instruction string, xwards bool) {
	steps, _ := strconv.Atoi(instruction[1:])
	if xwards {
		if string(instruction[0]) == "R" {
			facing.x = facing.y
		} else {
			facing.x = -facing.y
		}
		walk(steps, facing.x, coord, xwards)
		coord.x += facing.x * steps
	} else {
		if string(instruction[0]) == "R" {
			facing.y = -facing.x
		} else {
			facing.y = facing.x
		}
		walk(steps, facing.y, coord, xwards)
		coord.y += facing.y * steps
	}
}

func main() {
	input, err := os.Open("input/1")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	facing = point{x: 0, y: 1}
	coord = point{x: 0, y: 0}
	twice = point{x: 0, y: 0}
	origo = point{x: 0, y: 0}
	visited = make(map[point]int)
	xwards := true

	for scanner.Scan() {
		line := scanner.Text()
		// line := "R8, R4, R4, R8"
		instructions := strings.Split(line, ", ")
		for _, instruction := range instructions {
			newCoord(instruction, xwards)
			xwards = !xwards
		}
	}
	fmt.Printf("If we walk like normal we have to walk %d blocks\n", abs(coord.x)+abs(coord.y))
	fmt.Printf("If we read the back we have to walk %d blocks\n", abs(twice.x)+abs(twice.y))
}
