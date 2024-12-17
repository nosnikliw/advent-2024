package day17

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type MachineState struct {
	A int
	B int
	C int

	Program            []int
	InstructionPointer int

	Output []int
}

func LoadFile(fileName string) MachineState {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	a, _ := strconv.Atoi(strings.Fields(fileScanner.Text())[2])
	fileScanner.Scan()
	b, _ := strconv.Atoi(strings.Fields(fileScanner.Text())[2])
	fileScanner.Scan()
	c, _ := strconv.Atoi(strings.Fields(fileScanner.Text())[2])
	fileScanner.Scan()
	fileScanner.Scan()
	instructions := strings.Split(strings.Fields(fileScanner.Text())[1], ",")

	program := make([]int, len(instructions))
	for i, inst := range instructions {
		op, _ := strconv.Atoi(inst)
		program[i] = op
	}
	return MachineState{
		A:                  a,
		B:                  b,
		C:                  c,
		InstructionPointer: 0,
		Program:            program,
		Output:             []int{},
	}
}

func (machine MachineState) DisplayProgram() {
	for i := 0; i < len(machine.Program); i += 2 {
		fmt.Println(opNames[machine.Program[i]], machine.Program[i+1])
	}
}

func (machine MachineState) Run() []int {
	for machine.InstructionPointer >= 0 && machine.InstructionPointer < len(machine.Program) {
		operations[machine.Program[machine.InstructionPointer]](&machine, machine.Program[machine.InstructionPointer+1])
	}
	return machine.Output
}

func (machine MachineState) Expect(expected []int) bool {
	for machine.InstructionPointer >= 0 && machine.InstructionPointer < len(machine.Program) {
		operations[machine.Program[machine.InstructionPointer]](&machine, machine.Program[machine.InstructionPointer+1])
		if len(machine.Output) > len(expected) {
			return false
		}
		if len(machine.Output) > 0 {
			val := machine.Output[len(machine.Output)-1]
			if val != expected[len(machine.Output)-1] {
				return false
			}
		}
	}
	match := len(machine.Output) == len(expected)
	if match {
		fmt.Println("A:", machine.A)
		fmt.Println("B:", machine.B)
		fmt.Println("C:", machine.C)
		fmt.Println("Output:", machine.Output)
	}
	return match
}

var operations []Operation = []Operation{
	adv,
	bxl,
	bst,
	jnz,
	bxc,
	out,
	bdv,
	cdv,
}

var opNames []string = []string{
	"adv",
	"bxl",
	"bst",
	"jnz",
	"bxc",
	"out",
	"bdv",
	"cdv",
}

func comboValue(machine MachineState, operand int) int {
	if operand == 4 {
		// fmt.Println("Val:", machine.A)
		return machine.A
	} else if operand == 5 {
		return machine.B
	} else if operand == 6 {
		return machine.C
	} else if operand >= 0 && operand <= 3 {
		return operand
	}
	panic("bad operand")
}

type Operation func(*MachineState, int)

func adv(machine *MachineState, operand int) {
	// fmt.Println("adv", operand)
	(*machine).A = (*machine).A / int(math.Pow(2, float64(comboValue((*machine), operand))))
	(*machine).InstructionPointer += 2
}

func bxl(machine *MachineState, operand int) {
	// fmt.Println("bxl", operand)
	(*machine).B = (*machine).B ^ operand
	(*machine).InstructionPointer += 2
}

func bst(machine *MachineState, operand int) {
	// fmt.Println("bst", operand)
	(*machine).B = comboValue((*machine), operand) % 8
	(*machine).InstructionPointer += 2
}

func jnz(machine *MachineState, operand int) {
	// fmt.Println("jnz")
	if (*machine).A == 0 {
		(*machine).InstructionPointer += 2
	} else {
		(*machine).InstructionPointer = operand
	}
}

func bxc(machine *MachineState, _ int) {
	// fmt.Println("bxc")
	(*machine).B = (*machine).B ^ (*machine).C
	(*machine).InstructionPointer += 2
}

func out(machine *MachineState, operand int) {
	// fmt.Println("out", operand)
	(*machine).Output = append((*machine).Output, comboValue(*machine, operand)%8)
	(*machine).InstructionPointer += 2
}

func bdv(machine *MachineState, operand int) {
	// fmt.Println("bdv", operand)
	(*machine).B = (*machine).A / int(math.Pow(2, float64(comboValue((*machine), operand))))
	(*machine).InstructionPointer += 2
}

func cdv(machine *MachineState, operand int) {
	// fmt.Println("cxl", operand)
	(*machine).C = (*machine).A / int(math.Pow(2, float64(comboValue((*machine), operand))))
	(*machine).InstructionPointer += 2
}
