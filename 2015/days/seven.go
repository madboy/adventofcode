package days

import (
	"bufio"
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
	override := map[string]uint{"b": 46065}
	return fmt.Sprintf("%s\n%s", aSignal(lines, nil), aSignal(lines, override))
}

func aSignal(lines []string, override map[string]uint) string {
	wires := make(map[string]uint)
	var line string
	for len(lines) > 0 {
		success := false
		line, lines = lines[0], lines[1:]
		inst := strings.Split(line, " ")
		if strings.Contains(line, "AND") {
			if w1, ok1 := getValue(inst[0], wires, override); ok1 {
				if w2, ok2 := getValue(inst[2], wires, override); ok2 {
					wires[inst[4]] = w1 & w2
					success = true
				}
			}
		} else if strings.Contains(line, "OR") {
			if w1, ok1 := getValue(inst[0], wires, override); ok1 {
				if w2, ok2 := getValue(inst[2], wires, override); ok2 {
					wires[inst[4]] = w1 | w2
					success = true
				}
			}
		} else if strings.Contains(line, "NOT") {
			if w1, ok := getValue(inst[1], wires, override); ok {
				wires[inst[3]] = 65535 - w1
				success = true
			}
		} else if strings.Contains(line, "LSHIFT") {
			if w1, ok1 := getValue(inst[0], wires, override); ok1 {
				if w2, ok2 := getValue(inst[2], wires, override); ok2 {
					wires[inst[4]] = w1 << w2
					success = true
				}
			}
		} else if strings.Contains(line, "RSHIFT") {
			if w1, ok1 := getValue(inst[0], wires, override); ok1 {
				if w2, ok2 := getValue(inst[2], wires, override); ok2 {
					wires[inst[4]] = w1 >> w2
					success = true
				}
			}
		} else {
			if v, ok := getValue(inst[0], wires, override); ok {
				wires[inst[2]] = v
				success = true
			}
		}
		if !success {
			lines = append(lines, line)
		}
	}

	return fmt.Sprintf("The value of wire a is: %d", wires["a"])
}

func getValue(part string, wires map[string]uint, overrides map[string]uint) (uint, bool) {
	if override, ok := overrides[part]; ok {
		return override, true
	}
	value, err := strconv.ParseUint(part, 10, 16)
	if err != nil {
		if value, ok := wires[part]; ok {
			return value, true
		}
		return 0, false
	}
	return uint(value), true
}
