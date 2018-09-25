package days

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// Instruction is from the manual
type Instruction struct {
	i string
	r string
	o int
}

func execute(instructions []Instruction, register map[string]int) {
	idx := 0
	for idx >= 0 && idx < len(instructions) {
		inst := instructions[idx]
		switch inst.i {
		case "hlf":
			register[inst.r] /= 2
			idx++
		case "tpl":
			register[inst.r] *= 3
			idx++
		case "inc":
			register[inst.r]++
			idx++
		case "jmp":
			idx += inst.o
		case "jie":
			if (register[inst.r] % 2) == 0 {
				idx += inst.o
			} else {
				idx++
			}
		case "jio":
			if register[inst.r] == 1 {
				idx += inst.o
			} else {
				idx++
			}
		}
	}
}

// Run23 in which we take a close a look at registers
func Run23(scanner *bufio.Scanner) string {
	var instructions []Instruction
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		instruction := Instruction{i: parts[0]}
		if len(parts) == 3 {
			instruction.r = strings.Trim(parts[1], ",")
			instruction.o, _ = strconv.Atoi(parts[2])
		} else if instruction.i == "jmp" {
			instruction.o, _ = strconv.Atoi(parts[1])
		} else {
			instruction.r = parts[1]
		}
		instructions = append(instructions, instruction)
	}

	registersPart1 := map[string]int{
		"a": 0,
		"b": 0,
	}
	execute(instructions, registersPart1)

	registersPart2 := map[string]int{
		"a": 1,
		"b": 0,
	}
	execute(instructions, registersPart2)
	return fmt.Sprintf("Value in register b is %d for (a = 0) and %d for (a = 1)", registersPart1["b"], registersPart2["b"])
}
