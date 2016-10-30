package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func pairsOfLetters(s string) bool {
	//It contains a pair of any two letters that appears
	// at least twice in the string without overlapping,
	// like xyxy (xy) or aabcdefgaa (aa),
	// but not like aaa (aa, but it overlaps).

	// not completely correct
	// cannot identify non overlapping repeat chars like "aaaa`"

	var pairs map[string]int
	pairs = make(map[string]int)
	first := string(s[0])
	peak := ""
	peak2 := ""
	for i, c := range s[1:] {
		cur := string(c)
		if i+2 == len(s) {
			peak = ""
		} else {
			peak = string(s[i+2])
		}
		if i+3 >= len(s) {
			peak2 = ""
		} else {
			peak2 = string(s[i+3])
		}
		// looking at a section of 4
		// [1 2 3 4]
		// 1+2 cannot be equal to 2+3
		// but it's ok if 1+2 = 3+4
		if ((first + cur) == (peak + peak2)) || ((first + cur) != (cur + peak)) {
			pairs[first+cur]++
		}
		first = cur
	}
	for _, v := range pairs {
		if v >= 2 {
			return true
		}
	}
	return false
}

func repeatLetters(s string) bool {
	first := string(s[0])
	second := string(s[1])
	for _, c := range s[2:] {
		cur := string(c)
		if first == cur {
			return true
		}
		first = second
		second = cur
	}
	return false
}

func main() {
	nice := 0
	input, err := os.Open("day5.input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if pairsOfLetters(line) {
			if repeatLetters(line) {
				nice++
			}
		}
	}
	fmt.Printf("In total %d are nice\n", nice)
}
