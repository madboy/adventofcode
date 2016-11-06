package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var wires map[string]int

func gate(gateType string, lvalue, rvalue int) int {
	return 0
}

func getValue(part string) (int, error) {
	value, err := strconv.Atoi(part)
	if err != nil {
		if value, ok := wires[part]; ok {
			return value, nil
		}
		return 0, errors.New("We have no input")
	}
	return value, nil
}

func applyGate(receiver string, parts []string) bool {
	lvalue, err := getValue(parts[0])
	if err != nil {
		return false
	}
	rvalue, err := getValue(parts[2])
	if err != nil {
		return false
	}
	switch parts[1] {
	case "AND":
		wires[receiver] = (lvalue & rvalue)
	case "OR":
		wires[receiver] = (lvalue | rvalue)
	case "LSHIFT":
		wires[receiver] = (lvalue << uint8(rvalue))
	case "RSHIFT":
		wires[receiver] = (lvalue >> uint8(rvalue))
	}
	return true
}

func main() {
	wires = make(map[string]int)
	// eh, we should do this in another way
	for i := 0; i < 150; i++ {
		input, err := os.Open("input/7")
		if err != nil {
			log.Fatal(err)
		}
		defer input.Close()

		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			instruction := strings.Split(line, "->")
			signal := strings.Trim(instruction[0], " ")
			signalParts := strings.Split(signal, " ")
			receiver := strings.Trim(instruction[1], " ")
			// Determine which type of signal we have based on instruction length
			switch sl := len(signalParts); sl {
			case 1:
				value, err := getValue(signalParts[0])
				if err == nil {
					// fmt.Printf("We are getting past %s\n", line)
					wires[receiver] = value
				}
			case 2:
				if val, err := getValue(signalParts[1]); err == nil {
					// fmt.Printf("We are getting past %s\n", line)
					wires[receiver] = (65535 - val)
				}
			case 3:
				applyGate(receiver, signalParts)
			}
		}
	}
	var keys []string
	for k := range wires {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: %d\n", k, wires[k])
	}
}
