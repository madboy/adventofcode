package days

import (
	"bufio"
	"fmt"
)

// Run20 in which we are counting presents
func Run20(scanner *bufio.Scanner) string {
	const input int = 33100000
	const l = int(input / 10)
	var house [l]int
	var house2 [l]int
	elfVisits := make(map[int]int)

	for i := 1; i < l; i++ {
		for j := i; j < l; j += i {
			house[j] += i * 10

			if elfVisits[i] < 50 {
				house2[j] += i * 11
				elfVisits[i]++
			}
		}
	}

	h1 := 0
	h2 := 0
	for i, h := range house {
		if h >= input && h1 == 0 {
			h1 = i
		}

		if house2[i] >= input && h2 == 0 {
			h2 = i
		}

		if h1 != 0 && h2 != 0 {
			break
		}
	}

	return fmt.Sprintf("House with a lot of presents is: %d, and when we don't have inifinite visits it's: %d", h1, h2)
}
