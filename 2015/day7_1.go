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

func applyGate(receiver string, parts []string) {
	lvalue, lerr := getValue(parts[0])
	rvalue, rerr := getValue(parts[2])
	if (lerr == nil) && (rerr == nil) {
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
	}
}

func getInstruction(line string) (string, []string) {
	instruction := strings.Split(line, "->")
	signal := strings.Trim(instruction[0], " ")
	signalParts := strings.Split(signal, " ")
	receiver := strings.Trim(instruction[1], " ")
	return receiver, signalParts
}

func main() {
	wires = make(map[string]int)
	// eh, we should do this in another way
	// have a queue for gates that are not complete yet?
	// and pop as we start getting values
	// but only stop when the queue is empty?
	for i := 0; i < 150; i++ {
		input, err := os.Open("input/7")
		if err != nil {
			log.Fatal(err)
		}
		defer input.Close()

		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			line := scanner.Text()
			receiver, signalParts := getInstruction(line)
			// Determine which type of signal we have based on instruction length
			switch sl := len(signalParts); sl {
			case 1: //Pure value feed
				value, err := getValue(signalParts[0])
				if err == nil {
					wires[receiver] = value
				}
			case 2: // NOT
				if val, err := getValue(signalParts[1]); err == nil {
					wires[receiver] = (65535 - val)
				}
			case 3: // Two value gates
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
	fmt.Printf("And for a the signal is %d\n", wires["a"])
}
