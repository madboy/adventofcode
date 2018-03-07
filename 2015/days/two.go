package days

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Run2 for the second day of christmas
func Run2(scanner *bufio.Scanner) string {
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return fmt.Sprintf("%s\n%s", paperArea(lines), ribbonLength(lines))
}

func paperArea(lines []string) string {
	totalArea := 0
	for _, line := range lines {
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
	return fmt.Sprintf("The total area of paper the elves need is: %d", totalArea)
}

func ribbonLength(lines []string) string {
	totalLength := 0
	for _, line := range lines {
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
	return fmt.Sprintf("The total lenght of ribbon the elves need is: %d", totalLength)
}
