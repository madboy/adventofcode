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

func logicGate(parts []string, wires map[string]int) int {
	lvalue := wires[parts[0]]
	switch parts[1] {
	case "AND":
		if rvalue, ok := wires[parts[2]]; ok {
			return (lvalue & rvalue)
		} else {
			return -1
		}
	case "OR":
		if rvalue, ok := wires[parts[2]]; ok {
			return (lvalue | rvalue)
		} else {
			return -1
		}
	case "LSHIFT":
		rtmp, _ := strconv.Atoi(parts[2])
		rvalue := uint8(rtmp)
		return (lvalue << uint8(rvalue))
	case "RSHIFT":
		rtmp, _ := strconv.Atoi(parts[2])
		rvalue := uint8(rtmp)
		return (lvalue >> uint8(rvalue))
	}
	return 0
}

func main() {
	var wires map[string]int
	wires = make(map[string]int)
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
			value, _ := strconv.Atoi(signalParts[0])
			wires[receiver] = value
		case 2:
			if val, ok := wires[signalParts[1]]; ok {
				// value := wires[signalParts[1]]
				wires[receiver] = (65535 - val)
			}
		case 3:
			if _, ok := wires[signalParts[0]]; ok {
				val := logicGate(signalParts, wires)
				if val != -1 {
					wires[receiver] = logicGate(signalParts, wires)
				}
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
