package days

import (
	"bufio"
	"fmt"

	"github.com/madboy/adventofcode/2015/tools"
	"github.com/madboy/combinatoric"
)

func sum(c []int) int {
	sum := 0
	for _, v := range c {
		sum += v
	}
	return sum
}

// Run17 liquidates our eggnog
func Run17(scanner *bufio.Scanner) string {
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	eggnog := 150

	containers := []int{50, 48, 45, 44, 42, 41, 35, 35, 33, 30, 24, 20, 18, 18, 16, 14, 13, 13, 6, 1}

	var combinations [][]int

	for i := range tools.Range(0, len(containers)) {
		combos := combinatoric.Combinations(containers, i)
		for _, combination := range combos {
			if sum(combination) == eggnog {
				combinations = append(combinations, combination)
			}
		}
	}

	minContainers := len(containers)
	containerDistribution := make(map[int]int)
	for _, c := range combinations {
		l := len(c)
		containerDistribution[l]++

		if l < minContainers {
			minContainers = l
		}
	}

	return fmt.Sprintf("We have %d ways of storing our eggnog\nAnd %d uses the minimum of %d containers.", len(combinations), containerDistribution[minContainers], minContainers)
}
