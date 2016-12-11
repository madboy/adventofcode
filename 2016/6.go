package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type col struct {
	c map[string]int
}

func (co *col) Add(l string) {
	if co.c == nil {
		co.c = make(map[string]int)
	}
	co.c[l]++
}

func main() {
	input, err := os.Open("input/6")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	message := [8]*col{&col{}, &col{}, &col{}, &col{}, &col{}, &col{}, &col{}, &col{}}
	ecVersion := ""
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			message[i].Add(string(c))
		}
	}
	for _, m := range message {
		min := 100000
		common := ""
		for char, count := range m.c {
			if count < min {
				common = char
				min = count
			}
		}
		ecVersion += common
	}
	fmt.Println(ecVersion)
}
