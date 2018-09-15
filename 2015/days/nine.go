package days

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/madboy/adventofcode/2015/tools"
	"github.com/madboy/combinatoric"
)

type distance struct {
	from string
	to   string
	d    int
}

// Run9 will have us travelling
func Run9(scanner *bufio.Scanner) string {
	distances := make(map[string]int)
	locations := tools.NewSet()
	for scanner.Scan() {
		line := scanner.Text()
		entry := strings.Split(line, " ")
		distance, err := strconv.Atoi(entry[4])
		if err != nil {
			log.Fatal(fmt.Sprintf("%s does not contain a distance\n", line))
		}
		distances[fmt.Sprintf("%s%s", entry[0], entry[2])] = distance
		distances[fmt.Sprintf("%s%s", entry[2], entry[0])] = distance
		locations.Add(entry[0])
		locations.Add(entry[2])
	}

	currentMin := 10000000
	currentMax := 0
	indices := tools.Range(0, len(locations.Values))

	for _, p := range combinatoric.Permutations(indices, len(locations.Values)) {
		tdist := 0
		for i := 0; i < len(indices)-1; i++ {
			tdist += distances[fmt.Sprintf("%s%s", locations.Values[p[i]], locations.Values[p[i+1]])]
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
