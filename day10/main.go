package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lukegriffith/aoc22"
)

type program struct {
	Cycle     int
	Position  int
	xRegister int
	SignalSum int
}

func (p *program) Increment() {
	var (
		signalStrength = func(cycle int, registerX int) int { return p.Cycle * p.xRegister }
		isDebug        = func(cycle int) bool {
			if cycle == 20 {
				return true
			} else if (cycle-20)%40 == 0 {
				return true
			} else {
				return false
			}
		}
	)

	p.Cycle++
	p.Position++
	if isDebug(p.Cycle) {
		p.SignalSum = p.SignalSum + signalStrength(p.Cycle, p.xRegister)
	}

	if p.Position == p.xRegister-1 || p.Position == p.xRegister || p.Position+1 == p.xRegister {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
	if p.Cycle%40 == 0 {
		fmt.Println()
		p.Position = 0
	}
}

func (p *program) AddToRegister(i int) {
	p.xRegister = p.xRegister + i
}

func day10(inputs []string) int {
	var (
		prog        program = program{0, 0, 1, 0}
		noopCommand         = "noop"
		addCommand          = "addx"
	)
	for _, command := range inputs {
		csplit := strings.Split(command, " ")
		opCode := csplit[0]
		switch opCode {
		case noopCommand:
			prog.Increment()
		case addCommand:
			instruction, err := strconv.Atoi(csplit[1])
			aoc22.CheckErr(err)
			prog.Increment()
			prog.Increment()
			prog.AddToRegister(instruction)
		}

	}
	return prog.SignalSum
}
func main() {
	inputs := aoc22.LoadInputs("input_test.txt")
	fmt.Println(day10(inputs))
}
