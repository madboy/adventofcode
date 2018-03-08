package days

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Run7 risks crossing our wires
func Run7(scanner *bufio.Scanner) string {
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return fmt.Sprintf("%s\n%s", aSignal(lines), bOverride(lines))
}

func aSignal(lines []string) string {
	var wires map[string]int

	wires = make(map[string]int)
	// eh, we should do this in another way
	// have a queue for gates that are not complete yet?
	// and pop as we start getting values
	// but only stop when the queue is empty?
	for i := 0; i < 150; i++ {
		for _, line := range lines {
			receiver, signalParts := getInstruction(line)
			// Determine which type of signal we have based on instruction length
			switch sl := len(signalParts); sl {
			case 1: //Pure value feed
				value, err := getValue(signalParts[0], wires)
				if err == nil {
					wires[receiver] = value
				}
			case 2: // NOT
				if val, err := getValue(signalParts[1], wires); err == nil {
					wires[receiver] = (65535 - val)
				}
			case 3: // Two value gates
				applyGate(receiver, signalParts, wires)
			}
		}
	}
	return fmt.Sprintf("For a the signal is %d", wires["a"])
}

func bOverride(lines []string) string {
	var wires map[string]int
	var queue map[string]struct{}
	wires = make(map[string]int)
	queue = make(map[string]struct{})

	// Run loop at least once to get the diagram read
	for first := true; first || len(queue) != 0; {
		readDiagram(lines, wires, queue)
		first = false
	}
	return fmt.Sprintf("The value of wire a is %d", wires["a"])
}

// This could be made a whole lot better
// Not read the file multiple times
// Not evaluate gates and wires that already have their value set
func readDiagram(lines []string, wires map[string]int, queue map[string]struct{}) {
	for _, line := range lines {
		receiver, signalParts := getInstruction(line)
		// Determine which type of signal we have based on instruction length
		switch sl := len(signalParts); sl {
		case 1: //Pure value feed
			value, err := getValue(signalParts[0], wires)
			if err == nil {
				wires[receiver] = value
				delete(queue, signalParts[0])
			} else {
				queue[signalParts[0]] = struct{}{}
			}
		case 2: // NOT
			if val, err := getValue(signalParts[1], wires); err == nil {
				wires[receiver] = (65535 - val)
				delete(queue, signalParts[1])
			} else {
				queue[signalParts[1]] = struct{}{}
			}
		case 3: // Two value gates
			lvalue, lerr := getValue(signalParts[0], wires)
			if lerr != nil {
				queue[signalParts[0]] = struct{}{}
			}
			rvalue, rerr := getValue(signalParts[2], wires)
			if rerr != nil {
				queue[signalParts[2]] = struct{}{}
			}
			// A gate only works if both wires have input
			if (lerr == nil) && (rerr == nil) {
				delete(queue, signalParts[0])
				delete(queue, signalParts[2])
				applyGatev2(receiver, lvalue, rvalue, signalParts[1], wires)
			}
		}
	}
}

func applyGate(receiver string, parts []string, wires map[string]int) {
	lvalue, lerr := getValue(parts[0], wires)
	rvalue, rerr := getValue(parts[2], wires)
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

func applyGatev2(receiver string, lvalue, rvalue int, gate string, wires map[string]int) {
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

func getValue(part string, wires map[string]int) (int, error) {
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
