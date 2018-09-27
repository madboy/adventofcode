package days

import (
	"bufio"
	"fmt"
	"regexp"
)

// Aunt contains what we know about auntie
type Aunt map[string]string

func isAunt(a Aunt, clues map[string]string) bool {
	candidate := false
	for k, v := range clues {
		if a[k] == v || a[k] == "" {
			candidate = true
		} else {
			candidate = false
			break
		}
	}
	return candidate
}

func isRealAunt(a Aunt, clues map[string]string) bool {
	for k, v := range clues {
		candidate := false
		switch k {
		case "cats", "trees":
			candidate = (a[k] > v || a[k] == "")
		case "pomeranians", "goldfish":
			candidate = (a[k] < v || a[k] == "")
		default:
			candidate = (a[k] == v || a[k] == "")
		}
		if !candidate {
			return false
		}
	}
	return true
}

// Run16 will help us make the best cookies
func Run16(scanner *bufio.Scanner) string {
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	var aunts []Aunt

	match := regexp.MustCompile(`Sue (\d+): (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)`)
	for _, a := range input {
		m := match.FindStringSubmatch(a)
		tmp := Aunt{"name": m[1], m[2]: m[3], m[4]: m[5], m[6]: m[7]}
		aunts = append(aunts, tmp)
	}

	clues := map[string]string{
		"children":    "3",
		"cats":        "7",
		"samoyeds":    "2",
		"pomeranians": "3",
		"akitas":      "0",
		"vizslas":     "0",
		"goldfish":    "5",
		"trees":       "3",
		"cars":        "2",
		"perfumes":    "1",
	}

	auntie := "unknown"
	realAuntie := "unknown"

	for _, a := range aunts {
		if isAunt(a, clues) {
			auntie = a["name"]
		}

		if isRealAunt(a, clues) {
			realAuntie = a["name"]
		}
	}

	return fmt.Sprintf("The aunt who sent me the kit was: %s\nThe real aunt was in fact: %s", auntie, realAuntie)
}
