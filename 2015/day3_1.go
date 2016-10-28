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

func main() {
	var travels map[string]int
	travels = make(map[string]int)
	n, e := 0, 0
	currentCoord := calculateCoord(n, e)
	travels[currentCoord]++
	input, err := ioutil.ReadFile("day3.input")
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range input {
		switch direction := string(c); direction {
		case "^":
			n++
		case "v":
			n--
		case ">":
			e++
		case "<":
			e--
		}
		currentCoord = calculateCoord(n, e)
		travels[currentCoord]++
	}
	fmt.Printf("Santa has visited %d houses\n", len(travels))
}
