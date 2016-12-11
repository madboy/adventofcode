package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func orderLetters(letters map[string]int) []string {
	return []string{"a", "b", "c"}
}

func main() {
	input, err := os.Open("input/4.test")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		var letters = make(map[string]int)
		pparts := strings.Split(line, "[") // checksum will need the ] removed
		parts := strings.Split(pparts[0], "-")
		room, checkSum := parts[:len(parts)-1], parts[len(parts)]
		for _, w := range room {
			for _, c := range w {
				letters[string(c)]++
			}
		}
		// oletters := orderLetters(letters)
		var filter = make(map[string][]string)
		filter = map[string][]string{
			"apa": []string{"apa", "lapa"},
		}
		fmt.Println(letters)
		fmt.Println(checkSum)
		fmt.Println(filter)
	}
}
