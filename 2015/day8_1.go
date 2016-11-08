package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	input, err := os.Open("input/8")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	totalStringCode := 0
	totalStringChar := 0
	totalNewChars := 0

	reg, err := regexp.Compile(`\\x[a-fA-F0-9]{2}`)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		totalStringCode += len(line)
		totalNewChars += len(fmt.Sprintf("%q", line))

		// \" \\ \x
		line = strings.Trim(line, "\"")
		line = strings.Replace(line, "\\\"", "1", -1)
		line = strings.Replace(line, "\\\\", "1", -1)
		// Replace hex chars
		line = reg.ReplaceAllString(line, "-")
		totalStringChar += len(line)
	}
	fmt.Printf("%d - %d = %d\n", totalStringCode, totalStringChar, totalStringCode-totalStringChar)
	fmt.Printf("%d - %d = %d\n", totalNewChars, totalStringCode, totalNewChars-totalStringCode)
}
