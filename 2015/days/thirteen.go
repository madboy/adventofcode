package days

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/madboy/adventofcode/2015/tools"
)

// Seating contains a seating arrangement
type Seating struct {
	Person string
	// Gain      bool
	Points    int
	Companion string
}

// Run13 is helping us with optimal placing
func Run13(scanner *bufio.Scanner) string {
	// input := []string{
	// 	"David would gain 46 happiness units by sitting next to Alice.",
	// 	"David would lose 7 happiness units by sitting next to Bob.",
	// 	"David would gain 41 happiness units by sitting next to Carol.",
	// 	"Carol would lose 62 happiness units by sitting next to Alice.",
	// 	"Carol would gain 60 happiness units by sitting next to Bob.",
	// 	"Carol would gain 55 happiness units by sitting next to David.",
	// 	"Alice would gain 54 happiness units by sitting next to Bob.",
	// 	"Alice would lose 79 happiness units by sitting next to Carol.",
	// 	"Alice would lose 2 happiness units by sitting next to David.",
	// 	"Bob would gain 83 happiness units by sitting next to Alice.",
	// 	"Bob would lose 7 happiness units by sitting next to Carol.",
	// 	"Bob would lose 63 happiness units by sitting next to David.",
	// }

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	people := tools.NewSet()
	var seatings []Seating

	for _, p := range input {
		parts := strings.Split(p, " ")
		person, gain, companion := parts[0], parts[2] == "gain", strings.Trim(parts[10], ".")
		points, err := strconv.Atoi(parts[3])
		if err != nil {
			log.Fatal("couldn't read points", err)
		}
		if !gain {
			points = -points
		}
		seatings = append(seatings, Seating{person, points, companion})
		people.Add(person)
	}
	length := len(people.Values)
	indexes := tools.Range(length)
	combinations := tools.Permutations(indexes)
	max := 0
	for _, c := range combinations {
		sum := 0
		placement := []int{c[length-1]}
		placement = append(placement, c...)
		placement = append(placement, c[0])
		for i := 1; i <= length; i++ {
			l := placement[i-1]
			r := placement[i+1]
			p := placement[i]
			left, right := people.Values[l], people.Values[r]
			for _, s := range seatings {
				if s.Person == people.Values[p] && (s.Companion == left || s.Companion == right) {
					sum += s.Points
				}
			}
		}
		if sum > max {
			max = sum
		}
	}
	fmt.Println(people.Values)
	return fmt.Sprintf("%d", max)
}
