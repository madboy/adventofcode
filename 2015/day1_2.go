package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	input, err := ioutil.ReadFile("day1.input")
	if err != nil {
		log.Fatal(err)
	}
	floor := 0
	pos := 0 //reason for using pos is if there's any non () in the input
	for _, c := range input {
		if string(c) == "(" {
			floor++
			pos++
		} else if string(c) == ")" {
			floor--
			pos++
		}
		if floor == -1 {
			fmt.Printf("Position for entering basement is: %d\n", pos)
			break
		}
	}
}
