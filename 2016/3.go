package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isTriangle(sides [3]int) bool {
	sum1 := sides[0] + sides[1]
	sum2 := sides[1] + sides[2]
	sum3 := sides[0] + sides[2]
	return (sum1 > sides[2]) && (sum2 > sides[0]) && (sum3 > sides[1])
}

func readSides(lengths []string) [3]int {
	var sides [3]int
	sides[0], _ = strconv.Atoi(lengths[0])
	sides[1], _ = strconv.Atoi(lengths[1])
	sides[2], _ = strconv.Atoi(lengths[2])
	return sides
}

func main() {
	input, err := os.Open("input/3")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	rvalids := 0

	var ct1 [3]int
	var ct2 [3]int
	var ct3 [3]int
	col := 0
	cvalids := 0
	for scanner.Scan() {
		line := scanner.Text()
		lengths := strings.Fields(line)

		// Check triangle specified by row
		rsides := readSides(lengths)
		if isTriangle(rsides) {
			rvalids++
		}

		// Deal with column specified triangles
		if col == 2 {
			ct1[col], _ = strconv.Atoi(lengths[0])
			ct2[col], _ = strconv.Atoi(lengths[1])
			ct3[col], _ = strconv.Atoi(lengths[2])
			fmt.Println(ct1)
			if isTriangle(ct1) {
				cvalids++
			}
			if isTriangle(ct2) {
				cvalids++
			}
			if isTriangle(ct3) {
				cvalids++
			}
			ct1 = [3]int{0, 0, 0}
			ct2 = [3]int{0, 0, 0}
			ct3 = [3]int{0, 0, 0}
			col = 0
		} else {
			ct1[col], _ = strconv.Atoi(lengths[0])
			ct2[col], _ = strconv.Atoi(lengths[1])
			ct3[col], _ = strconv.Atoi(lengths[2])
			col++
		}
	}
	fmt.Printf("There's %d valid triangles listed if you read by rows\n", rvalids)
	fmt.Printf("There's %d valid triangles listed if you read by columns\n", cvalids)
}
