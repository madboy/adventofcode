package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("input/1")
	if err != nil {
		log.Fatal(err)
	}
	floor := 0
	for _, c := range input {
		if string(c) == "(" {
			floor++
		}
		if string(c) == ")" {
			floor--
		}
	}
	fmt.Printf("Santa ends up on floor: %d\n", floor)
}
