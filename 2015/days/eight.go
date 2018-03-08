package days

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"
)

// Run8 will have us be in character
func Run8(scanner *bufio.Scanner) string {
	totalStringCode := 0
	totalStringChar := 0
	totalNewChars := 0

	reg, err := regexp.Compile(`\\x[a-fA-F0-9]{2}`)
	if err != nil {
		log.Fatal(err)
	}

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
	return fmt.Sprintf("%d - %d = %d\n%d - %d = %d",
		totalStringCode, totalStringChar, totalStringCode-totalStringChar,
		totalNewChars, totalStringCode, totalNewChars-totalStringCode)
}
