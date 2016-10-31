package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func updateLights(action string, from []string, to []string, lights map[string]int) {
	fromX, _ := strconv.Atoi(from[0])
	fromY, _ := strconv.Atoi(from[1])
	toX, _ := strconv.Atoi(to[0])
	toY, _ := strconv.Atoi(to[1])
	for x := fromX; x <= toX; x++ {
		for y := fromY; y <= toY; y++ {
			switch action {
			case "on":
				lights[strconv.Itoa(x)+"x"+strconv.Itoa(y)] = 1
			case "off":
				lights[strconv.Itoa(x)+"x"+strconv.Itoa(y)] = 0
			case "toggle":
				if lights[strconv.Itoa(x)+"x"+strconv.Itoa(y)] == 0 {
					lights[strconv.Itoa(x)+"x"+strconv.Itoa(y)] = 1
				} else {
					lights[strconv.Itoa(x)+"x"+strconv.Itoa(y)] = 0
				}
			}
		}
	}
}

func main() {
	var lights map[string]int
	lights = make(map[string]int)
	lit := 0
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
	fmt.Printf("We have %d lights lit\n", lit)
}
