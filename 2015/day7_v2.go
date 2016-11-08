package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
func getInstruction(line string) (string, []string) {
	instruction := strings.Split(line, "->")
	signal := strings.Trim(instruction[0], " ")
	signalParts := strings.Split(signal, " ")
	receiver := strings.Trim(instruction[1], " ")
	return receiver, signalParts
}

func applyGate(receiver string, lvalue, rvalue int, gate string) {
	switch gate {
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

// This could be made a whole lot better
// Not read the file multiple times
// Not evaluate gates and wires that already have their value set
func readDiagram() {
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
				delete(queue, signalParts[0])
			} else {
				queue[signalParts[0]] = struct{}{}
			}
		case 2: // NOT
			if val, err := getValue(signalParts[1]); err == nil {
				wires[receiver] = (65535 - val)
				delete(queue, signalParts[1])
			} else {
				queue[signalParts[1]] = struct{}{}
			}
		case 3: // Two value gates
			lvalue, lerr := getValue(signalParts[0])
			if lerr != nil {
				queue[signalParts[0]] = struct{}{}
			}
			rvalue, rerr := getValue(signalParts[2])
			if rerr != nil {
				queue[signalParts[2]] = struct{}{}
			}
			// A gate only works if both wires have input
			if (lerr == nil) && (rerr == nil) {
				delete(queue, signalParts[0])
				delete(queue, signalParts[2])
				applyGate(receiver, lvalue, rvalue, signalParts[1])
			}
		}
	}
}

var wires map[string]int
var queue map[string]struct{}

func main() {
	wires = make(map[string]int)
	queue = make(map[string]struct{})

	// Run loop at least once to get the diagram read
	for first := true; first || len(queue) != 0; {
		readDiagram()
		first = false
	}
	fmt.Println(queue)
	fmt.Println(wires)
	fmt.Printf("The value of wire a is %d\n", wires["a"])
}
