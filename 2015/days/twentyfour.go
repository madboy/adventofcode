package days

import (
	"bufio"
	"fmt"
	"log"
	"strconv"

	"github.com/madboy/adventofcode/2015/tools"
	"github.com/madboy/combinatoric"
)

func removeElement(e int, arr []int) []int {
	var tmp []int
	for i, v := range arr {
		if v == e {
			tmp = append(tmp, arr[i+1:]...)
			return tmp
		}
		tmp = append(tmp, v)
	}
	return tmp
}

func removeElements(from, arr []int) []int {
	tmp := from
	for _, e := range arr {
		tmp = removeElement(e, tmp)
	}
	return tmp
}

// get all combinations of packages that match the taget weight
func possibleGroups(tw, nbr int, arr []int) [][]int {
	var groups [][]int
	for _, c := range combinatoric.Combinations(arr, nbr) {
		if tools.SumInt(c) == tw {
			groups = append(groups, c)
		}
	}
	return groups
}

// quantum engtanglement is the product of all package weights in the group
func calculateQE(arr []int) int {
	qe := 1
	for _, v := range arr {
		qe *= v
	}
	return qe
}

func quantumEntanglement(tw int, packages []int) int {
	// Get all possible groups of packages starting with
	// a small a groups as possible, ie the passenger compartment
	// as we go through the groups remove the packages we have
	// from the list of possible packages
	// Start creating new possible groups for the left side of the sleigh
	// based on the trimmed list of packages
	// Remove those from the shorter list and see if we can create any
	// third group (right side of sleigh) that matches requirements
	// If we can we are done* and can calculate the QE
	// * since the packages are sorted and we start small the first match
	// we have should have the smallest possible quantum entanglement
	for i := 2; i < len(packages); i++ {
		pGroups := possibleGroups(tw, i, packages)

		for _, candidate := range pGroups {
			tmp := removeElements(packages, candidate)

			for j := i; j < len(tmp)-i; j++ {
				leftSideGroups := possibleGroups(tw, j, tmp)

				for _, lCandidate := range leftSideGroups {
					rTmp := removeElements(tmp, lCandidate)
					rightSideGroups := possibleGroups(tw, len(rTmp), rTmp)
					if len(rightSideGroups) > 0 {
						return calculateQE(candidate)
					}
				}
			}
		}
	}
	return -1
}

func quantumEntanglementWithTrunk(tw int, packages []int) int {
	// Same as for the other just adding another group
	for i := 2; i < len(packages); i++ {
		pGroups := possibleGroups(tw, i, packages)

		for _, candidate := range pGroups {
			tmp := removeElements(packages, candidate)

			for j := i; j < len(tmp)-i; j++ {
				leftSideGroups := possibleGroups(tw, j, tmp)

				for _, lCandidate := range leftSideGroups {
					rTmp := removeElements(tmp, lCandidate)

					for k := j; k < len(rTmp)-j; k++ {
						rightSideGroups := possibleGroups(tw, k, rTmp)

						for _, rCandidate := range rightSideGroups {
							trunkTmp := removeElements(rTmp, rCandidate)
							trunkGroups := possibleGroups(tw, len(trunkTmp), trunkTmp)
							if len(trunkGroups) > 0 {
								return calculateQE(candidate)
							}
						}
					}
				}
			}
		}
	}
	return -1
}

// Run24 in which we balance the sleigh
func Run24(scanner *bufio.Scanner) string {
	var input []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal("not a number")
		}
		input = append(input, i)
	}
	totalW := tools.SumInt(input)
	targetW := totalW / 3
	targetWtrunk := totalW / 4

	return fmt.Sprintf("Entaglement of the sleigh is: %d\nEntaglement of the sleigh (incl trunk) is: %d",
		quantumEntanglement(targetW, input),
		quantumEntanglementWithTrunk(targetWtrunk, input))
}
