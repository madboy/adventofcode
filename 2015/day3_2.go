package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func calculateCoord(n, e int) string {
	tn := strconv.Itoa(n)
	te := strconv.Itoa(e)
	return strings.Join([]string{tn, te}, ",")
}

func calculateMovement(input byte, n, e int) (int, int) {
	switch direction := string(input); direction {
	case "^":
		n++
	case "v":
		n--
	case ">":
		e++
	case "<":
		e--
	}
	return n, e
}

func main() {
	var travels map[string]int
	travels = make(map[string]int)
	sn, se, rn, re := 0, 0, 0, 0

	travels[calculateCoord(sn, se)] += 2
	input, err := ioutil.ReadFile("day3.input")
	if err != nil {
		log.Fatal(err)
	}
	for i, c := range input {
		if (i % 2) == 0 {
			sn, se = calculateMovement(c, sn, se)
			travels[calculateCoord(sn, se)]++
		} else {
			rn, re = calculateMovement(c, rn, re)
			travels[calculateCoord(rn, re)]++
		}
	}
	fmt.Printf("Santa and Robo-Santa have visted %d houses\n", len(travels))
}
