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

func printSeating(finalSeating []int, names []string) {
	for _, s := range finalSeating[0 : len(finalSeating)-1] {
		fmt.Printf("%s\t", names[s])
	}
	fmt.Println()
}

func getKey(name1, name2 string) string {
	if name1[0] < name2[0] {
		return fmt.Sprintf("%s%s", name1, name2)
	}
	return fmt.Sprintf("%s%s", name2, name1)
}

func getSeating(input []string) (tools.Set, map[string]int) {
	people := tools.NewSet()
	keys := make(map[string]int)

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
		keys[getKey(person, companion)] += points
		people.Add(person)
	}
	return people, keys
}

func addMeToSeating(people tools.Set, keys map[string]int) (tools.Set, map[string]int) {
	for _, p := range people.Values {
		keys[getKey("Me", p)] += 0
	}
	people.Add("Me")
	return people, keys
}

func getMaxHappiness(people tools.Set, seatingKeys map[string]int) string {
	length := len(people.Values)
	indexes := tools.Range(0, length)
	combinations := combinatoric.Permutations(indexes, length)
	max := 0
	var finalSeating []int
	for _, c := range combinations {
		sum := 0
		placement := c
		placement = append(placement, c[0])
		for i := 0; i < length; i++ {
			left, right := people.Values[placement[i]], people.Values[placement[i+1]]
			key := getKey(left, right)
			sum += seatingKeys[key]
		}

		if sum > max {
			max = sum
			finalSeating = placement
		}
	}

	printSeating(finalSeating, people.Values)
	return fmt.Sprintf("%d", max)
}

// Run13 is helping us with optimal placing
func Run13(scanner *bufio.Scanner) string {
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	people, seatingKeys := getSeating(input)
	part1 := getMaxHappiness(people, seatingKeys)
	people, seatingKeys = addMeToSeating(people, seatingKeys)
	part2 := getMaxHappiness(people, seatingKeys)
	return fmt.Sprintf("Seating happiness: %s, Seating happiness (incl me): %s", part1, part2)
}
