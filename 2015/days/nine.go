package days

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/madboy/adventofcode/2015/tools"
)

type distance struct {
	from string
	to   string
	d    int
}

// Run9 will have us travelling
func Run9(scanner *bufio.Scanner) string {
	distances := make(map[string]int)
	locations := make(map[string]bool)
	for scanner.Scan() {
		line := scanner.Text()
		entry := strings.Split(line, " ")
		distance, err := strconv.Atoi(entry[4])
		if err != nil {
			log.Fatal(fmt.Sprintf("%s does not contain a distance\n", line))
		}
		distances[fmt.Sprintf("%s%s", entry[0], entry[2])] = distance
		distances[fmt.Sprintf("%s%s", entry[2], entry[0])] = distance
		locations[entry[0]] = true
		locations[entry[2]] = true
	}

	currentMin := 10000000
	currentMax := 0
	var keys []string
	for k := range locations {
		keys = append(keys, k)
	}
	indices := tools.Range(len(keys))

	for _, p := range tools.Permutations(indices) {
		tdist := 0
		for i := 0; i < len(indices)-1; i++ {
			tdist += distances[fmt.Sprintf("%s%s", keys[p[i]], keys[p[i+1]])]
		}
		if tdist < currentMin {
			currentMin = tdist
		}
		if tdist > currentMax {
			currentMax = tdist
		}
	}
	return fmt.Sprintf("shortest distance: %d, longest distance: %d", currentMin, currentMax)
}
