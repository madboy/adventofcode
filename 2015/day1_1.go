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
		switch char := string(c); char {
		case "(":
			floor++
		case ")":
			floor--
		}
	}
	fmt.Printf("Santa ends up on floor: %d\n", floor)
}
