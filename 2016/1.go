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

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	input, err := os.Open("input/1")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	facing := point{x: 0, y: 1}
	coord := point{x: 0, y: 0}
	twice := point{x: 0, y: 0}
	origo := point{x: 0, y: 0}
	var visited = make(map[point]int)
	xwards := true

	for scanner.Scan() {
		line := scanner.Text()
		// line := "R8, R4, R4, R8"
		instructions := strings.Split(line, ", ")
		for _, instruction := range instructions {
			steps, _ := strconv.Atoi(instruction[1:])
			if xwards {
				if string(instruction[0]) == "R" {
					facing.x = facing.y
				} else {
					facing.x = -facing.y
				}
				startx := coord.x
				for i := 0; i < steps; i++ {
					startx += facing.x * 1
					p := point{x: startx, y: coord.y}
					visited[p]++
					if visited[p] > 1 && (twice == origo) {
						twice = p
					}
				}
				coord.x += facing.x * steps
			} else {
				if string(instruction[0]) == "R" {
					facing.y = -facing.x
				} else {
					facing.y = facing.x
				}
				starty := coord.y
				for i := 0; i < steps; i++ {
					starty += facing.y * 1
					p := point{x: coord.x, y: starty}
					visited[p]++
					if visited[p] > 1 && (twice == origo) {
						twice = p
					}
				}
				coord.y += facing.y * steps
			}
			xwards = !xwards
		}
	}
	fmt.Printf("If we walk like normal we have to walk %d blocks\n", abs(coord.x)+abs(coord.y))
	fmt.Printf("If we read the back we have to walk %d blocks\n", abs(twice.x)+abs(twice.y))
}
