package days

import (
	"bufio"
	"fmt"

	"github.com/madboy/adventofcode/2015/tools"
)

// Run10 is having us looing-and-saying numbers
func Run10(scanner *bufio.Scanner) string {
	input := "1113122113"
	for _ = range tools.Range(40) {
		input = lookAndSay(input)
	}
	return fmt.Sprintf("Length of number is: %d", len(input))
}

func lookAndSay(number string) string {
	// adding " " at the end ensures we'll always end on mismatch
	bnumber := []byte(number + " ")
	current := bnumber[0]
	count := 1
	newNumber := ""
	for _, b := range bnumber[1:] {
		if current == b {
			count++
		} else {
			newNumber += fmt.Sprintf("%d%s", count, string(current))
			current = b
			count = 1
		}
	}
	return newNumber
}
