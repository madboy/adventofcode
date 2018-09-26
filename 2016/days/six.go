package days

import (
	"bufio"
	"fmt"
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

// Run6 in which we are reading Santas message
func Run6(scanner *bufio.Scanner) string {
	message := [8]*col{&col{}, &col{}, &col{}, &col{}, &col{}, &col{}, &col{}, &col{}}
	ecVersion := ""
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
	return fmt.Sprintf("%s", ecVersion)
}
