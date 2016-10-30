package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func countVowels(s string) int {
	sum := 0
	for _, c := range "aeiou" {
		sum += strings.Count(s, string(c))
	}
	return sum
}

func approvedString(s string) bool {
	forbidden := []string{"ab", "cd", "pq", "xy"}
	for _, f := range forbidden {
		if strings.Contains(s, f) {
			return false
		}
	}
	return true
}

func hasDoubleLetters(s string) bool {
	prev := string(s[0])
	cur := ""
	for _, c := range s[1:] {
		cur = string(c)
		if prev == string(cur) {
			return true
		}
		prev = cur
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
		if approvedString(line) {
			vowelCount := countVowels(line)
			if hasDoubleLetters(line) {
				if vowelCount >= 3 {
					nice++
				}
			}
		}
	}
	fmt.Printf("In total %d are nice\n", nice)
}
