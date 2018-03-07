package days

import (
	"bufio"
	"fmt"
	"strings"
)

// Run5 is the naught or nice list
func Run5(scanner *bufio.Scanner) string {
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return fmt.Sprintf("%s\n%s", niceOne(lines), niceTwo(lines))
}

func niceOne(lines []string) string {
	nice := 0
	for _, line := range lines {
		if approvedString(line) {
			vowelCount := countVowels(line)
			if hasDoubleLetters(line) {
				if vowelCount >= 3 {
					nice++
				}
			}
		}
	}
	return fmt.Sprintf("In total %d are nice", nice)
}

func niceTwo(lines []string) string {
	nice := 0
	for _, line := range lines {
		if pairsOfLetters(line) {
			if repeatLetters(line) {
				nice++
			}
		}
	}
	return fmt.Sprintf("In total %d are nice", nice)
}

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
