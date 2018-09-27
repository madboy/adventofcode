package days

import (
	"bufio"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/madboy/adventofcode/2015/tools"
)

type replacement struct {
	m string
	r string
}

// Shuffle randomizes order of elments in array
func shuffle(arr []replacement) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	l := len(arr) - 1
	for i := 0; i <= l; i++ {
		n := r1.Intn(l)
		swap(arr, i, n)
	}
}

func swap(arr []replacement, f, t int) {
	arr[f], arr[t] = arr[t], arr[f]
}

func moleculeVariations(medicineMolecule string, replacements []replacement) int {
	molecules := tools.NewSet()
	for _, r := range replacements {
		for k := range medicineMolecule {
			if k+len(r.m) < len(medicineMolecule) && medicineMolecule[k:k+len(r.m)] == r.m {
				variant := medicineMolecule[:k] + r.r + medicineMolecule[k+len(r.m):]
				molecules.Add(variant)
			}
		}
	}
	return len(molecules.Values)
}

// part 2 is from https://www.reddit.com/r/adventofcode/comments/3xflz8/day_19_solutions/cy4cu5b
func moleculeFabrication(medicineMolecule string, replacements []replacement) int {
	candidate := medicineMolecule
	count := 0
	for candidate != "e" {
		startMolecule := candidate
		for _, r := range replacements {
			if !strings.Contains(candidate, r.r) {
				continue
			}

			// make replacement of molecules in the opposite order
			// going from a replacement to the base molecule
			candidate = strings.Replace(candidate, r.r, r.m, 1)
			count++

			if startMolecule == candidate {
				candidate = medicineMolecule
				count = 0
				shuffle(replacements)
			}
		}
	}
	return count
}

// Run19 in which we go looking for medicine
func Run19(scanner *bufio.Scanner) string {
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	medicineMolecule := strings.Trim(input[len(input)-1], "")

	var replacements []replacement
	re := regexp.MustCompile(`(\w+) => (\w+)`)
	for _, l := range input[:len(input)-2] {
		matches := re.FindStringSubmatch(l)

		replacements = append(replacements, replacement{m: matches[1], r: matches[2]})
	}

	variations := moleculeVariations(medicineMolecule, replacements)
	count := moleculeFabrication(medicineMolecule, replacements)

	return fmt.Sprintf("Number of distinct molecules we have: %d, and it will take %d replacements to get to the medicine", variations, count)
}
