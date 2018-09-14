package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/madboy/adventofcode/2015/days"
)

type fn func(*bufio.Scanner) string

// Solution represents the code for a day in the calendar
type Solution struct {
	Function func(*bufio.Scanner) string
	Input    bool
}

var scanner *bufio.Scanner

func main() {
	day := flag.String("day", "1", "Day of Advent")

	var days = map[string]Solution{
		"1":  Solution{Function: days.Run1, Input: true},
		"2":  Solution{Function: days.Run2, Input: true},
		"3":  Solution{Function: days.Run3, Input: true},
		"4":  Solution{Function: days.Run4, Input: false},
		"5":  Solution{Function: days.Run5, Input: true},
		"6":  Solution{Function: days.Run6, Input: true},
		"7":  Solution{Function: days.Run7, Input: true},
		"8":  Solution{Function: days.Run8, Input: true},
		"9":  Solution{Function: days.Run9, Input: true},
		"10": Solution{Function: days.Run10, Input: false},
		"11": Solution{Function: days.Run11, Input: false},
		"12": Solution{Function: days.Run12, Input: true},
		"13": Solution{Function: days.Run13, Input: true},
		"14": Solution{Function: days.Run14, Input: true},
		"15": Solution{Function: days.Run15, Input: true},
		"16": Solution{Function: days.Run16, Input: true},
	}

	flag.Parse()
	if days[*day].Input {
		filename := "input/" + *day
		input, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer input.Close()
		scanner = bufio.NewScanner(input)
	}

	fmt.Println(days[*day].Function(scanner))
}
