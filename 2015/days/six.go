package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Run6 will give us more light than we need
func Run6(scanner *bufio.Scanner) string {
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return fmt.Sprintf("%s\n%s", lightsLit(lines), totalBrightness(lines))
}

func lightsLit(lines []string) string {
	var lights map[string]int
	lights = make(map[string]int)
	lit := 0

	for _, line := range lines {
		words := strings.Split(line, " ")
		// we have two cases
		// toggle, and
		// turn on or off
		// when we know case we extract coords
		if words[0] == "toggle" {
			from := strings.Split(words[1], ",")
			to := strings.Split(words[3], ",")
			updateLights("toggle", from, to, lights)
		} else {
			from := strings.Split(words[2], ",")
			to := strings.Split(words[4], ",")
			if words[1] == "on" {
				updateLights("on", from, to, lights)
			} else {
				updateLights("off", from, to, lights)
			}
		}
	}
	for _, v := range lights {
		if v == 1 {
			lit++
		}
	}
	return fmt.Sprintf("We have %d lights lit", lit)
}

func totalBrightness(lines []string) string {
	var lights map[string]int
	lights = make(map[string]int)
	lux := 0
	input, err := os.Open("input/6")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		// we have two cases
		// toggle, and
		// turn on or off
		// when we know case we extract coords
		//toggle: increase brightness of 2 of all selected lights
		if words[0] == "toggle" {
			from := strings.Split(words[1], ",")
			to := strings.Split(words[3], ",")
			updateLightsv2("toggle", from, to, lights)
			// turn on/off: increase or decrease brightness by 1
			// NB! can never go below 0
		} else {
			from := strings.Split(words[2], ",")
			to := strings.Split(words[4], ",")
			if words[1] == "on" {
				updateLightsv2("on", from, to, lights)
			} else {
				updateLightsv2("off", from, to, lights)
			}
		}
	}
	for _, v := range lights {
		lux += v
	}
	return fmt.Sprintf("We have %d total lux", lux)
}

func updateLights(action string, from []string, to []string, lights map[string]int) {
	fromX, _ := strconv.Atoi(from[0])
	fromY, _ := strconv.Atoi(from[1])
	toX, _ := strconv.Atoi(to[0])
	toY, _ := strconv.Atoi(to[1])
	for x := fromX; x <= toX; x++ {
		for y := fromY; y <= toY; y++ {
			coord := strconv.Itoa(x) + "x" + strconv.Itoa(y)
			switch action {
			case "on":
				lights[coord] = 1
			case "off":
				lights[coord] = 0
			case "toggle":
				lights[coord] = 1 - lights[coord]
			}
		}
	}
}

func updateLightsv2(action string, from []string, to []string, lights map[string]int) {
	fromX, _ := strconv.Atoi(from[0])
	fromY, _ := strconv.Atoi(from[1])
	toX, _ := strconv.Atoi(to[0])
	toY, _ := strconv.Atoi(to[1])
	for x := fromX; x <= toX; x++ {
		for y := fromY; y <= toY; y++ {
			coord := strconv.Itoa(x) + "x" + strconv.Itoa(y)
			switch action {
			case "on":
				lights[coord]++
			case "off":
				if lights[coord] > 0 {
					lights[coord]--
				} else {
					lights[coord] = 0
				}
			case "toggle":
				lights[coord] += 2
			}
		}
	}
}
