package days

import (
	"bufio"
	"fmt"
	"strings"
)

// Run11 will give us the most secure password. ever.
func Run11(scanner *bufio.Scanner) string {
	part1 := createValidPassword("hepxcrrq")
	part2 := createValidPassword(part1)
	return fmt.Sprintf("New password: %s, even newer password %s", part1, part2)
}

func createValidPassword(old string) string {
	new := increment(old)
	for !valid(new) {
		new = increment(new)
	}
	return new
}
func increment(old string) string {
	// 97 - 122 is the range of characters we're working with
	chars := []byte(old)
	current := len(chars) - 1
	for {
		if chars[current] < 122 {
			chars[current]++
			break
		} else {
			chars[current] = 97
			current--
		}
	}
	return string(chars)
}

func valid(password string) bool {
	return onlyAllowedCharacters(password) &&
		containsStraight(password) &&
		hasTwoPairs(password)
}

func containsStraight(password string) bool {
	// password should have charcters in ascending order
	// examples: abc, mno, ghi
	a := password[0]
	b := password[1]
	for _, c := range []byte(password[2:]) {
		if b-a == 1 && c-b == 1 {
			return true
		}
		a = b
		b = c
	}
	return false
}

func onlyAllowedCharacters(password string) bool {
	// password cannot contain i, o, l
	if strings.Contains(password, "i") || strings.Contains(password, "o") || strings.Contains(password, "l") {
		return false
	}
	return true
}

func hasTwoPairs(password string) bool {
	// password must have two different non overlapping pairs
	// example: bbabcoo
	p := password[0]
	pair := ""
	for _, c := range []byte(password[1:]) {
		// as pairs aren't allowed to overlap add first match to pair
		// so that we have something to check against if we get a second match
		if p == c && len(pair) < 1 {
			pair += string(p)
		} else if p == c && string(p) != pair {
			return true
		}
		p = c
	}
	return false
}
