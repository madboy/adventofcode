package days

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// Run12 is helping us do the accounting
func Run12(scanner *bufio.Scanner) string {
	cmd := exec.Command("./days/twelve.py")
	input := ""
	for scanner.Scan() {
		input += scanner.Text()
	}
	cmd.Stdin = strings.NewReader(input)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", out.String())
}
