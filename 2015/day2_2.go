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
	input, err := os.Open("input/2")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	totalLength := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		sides := strings.Split(line, "x")
		w, _ := strconv.Atoi(sides[0])
		h, _ := strconv.Atoi(sides[1])
		l, _ := strconv.Atoi(sides[2])
		front := (w + h) * 2
		side := (h + l) * 2
		top := (l + w) * 2
		volume := w * h * l
		length := int(math.Min(math.Min(float64(front), float64(side)), float64(top))) + volume
		totalLength += length
	}
	fmt.Printf("The total lenght of ribbon the elves need is: %d\n", totalLength)
}
