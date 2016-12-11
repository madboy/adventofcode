package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Letter struct {
	Char  string
	Count int
}

func (l Letter) String() string {
	return fmt.Sprintf("%s: %d", l.Char, l.Count)
}

type ByCount []Letter

func (a ByCount) Len() int      { return len(a) }
func (a ByCount) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool {
	if a[i].Count == a[j].Count {
		return a[i].Char < a[j].Char
	} else {
		return a[i].Count > a[j].Count
	}
}

func updateLetters(letters *[]Letter, letter string) {
	match := -1
	for i, l := range *letters {
		if l.Char == letter {
			match = i
		}
	}
	if match != -1 {
		(*letters)[match].Count++
	} else {
		*letters = append(*letters, Letter{Char: letter, Count: 1})
	}
}

func getRoomName(room []string, r int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var alphaMap = map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
		"i": 8,
		"j": 9,
		"k": 10,
		"l": 11,
		"m": 12,
		"n": 13,
		"o": 14,
		"p": 15,
		"q": 16,
		"r": 17,
		"s": 18,
		"t": 19,
		"u": 20,
		"v": 21,
		"w": 22,
		"x": 23,
		"y": 24,
		"z": 25,
	}
	name := ""
	for _, w := range room {
		for _, c := range w {
			i := alphaMap[string(c)]
			ii := (i + r) % 26
			name += string(alphabet[ii])
		}
		name += " "
	}
	return name
}

func getCheckSum(letters []Letter) string {
	checkSum := ""
	for _, c := range letters {
		checkSum += c.Char
	}
	return checkSum[:5]
}

func main() {
	input, err := os.Open("input/4")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	sum := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		// var letters = make(map[string]int)
		letters := []Letter{}
		pparts := strings.Split(line, "[") // checksum will need the ] removed
		parts := strings.Split(pparts[0], "-")
		room, sector, checkSum := parts[:len(parts)-1], parts[len(parts)-1], pparts[len(pparts)-1]
		checkSum = checkSum[:len(checkSum)-1]
		for _, w := range room {
			for _, c := range w {
				updateLetters(&letters, string(c))
			}
		}
		sort.Sort(ByCount(letters))
		cs := getCheckSum(letters)
		if cs == checkSum {
			n, _ := strconv.Atoi(sector)
			sum += n

			rot := (n % 26)
			fmt.Printf("Room name: %s, and sector id: %d\n", getRoomName(room, rot), n)
		}
	}
	fmt.Printf("The sum of the approved rooms is: %d\n", sum)
}
