package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("day2.input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	totalArea := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		sides := strings.Split(line, "x")
		w, _ := strconv.Atoi(sides[0])
		h, _ := strconv.Atoi(sides[1])
		l, _ := strconv.Atoi(sides[2])
		front := w * h
		side := h * l
		top := l * w
		area := 2*front + 2*side + 2*top + int(math.Min(math.Min(float64(front), float64(side)), float64(top)))
		totalArea += area
	}
	fmt.Printf("The total area of paper the elves need is: %d\n", totalArea)
}
