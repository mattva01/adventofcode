package compiler

import (
	"errors"
	"strconv"
	"strings"
)

type Computer struct {
	ProgramCounter int
	Instructions   []Instruction
	Accumulator    int
}

type Instruction struct {
	Code string
	Arg  int
}

type TrackedComputer struct {
	*Computer
	Visited map[int]struct{}
}

func NewComputer() *Computer {
	return &Computer{
		Accumulator:    0,
		ProgramCounter: 0,
		Instructions:   []Instruction{},
	}
}

func NewTrackedComputer() *TrackedComputer {
	return &TrackedComputer{
		Computer: NewComputer(),
		Visited:  make(map[int]struct{}),
	}
}

var EndOfProgram = errors.New("End of Program")

var ErrHitInfiniteLoop = errors.New("Infinite loop detected")

func (c *Computer) LoadProgram(program []string) error {
	var instructions []Instruction
	for _, line := range program {
		lineparts := strings.Split(line, " ")
		arg, err := strconv.Atoi(lineparts[1])
		if err != nil {
			return err
		}
		instruction := Instruction{
			Code: lineparts[0],
			Arg:  arg,
		}
		instructions = append(instructions, instruction)
	}
	c.Instructions = instructions
	return nil
}

func (c *Computer) ExecuteInstruction() error {
	if c.ProgramCounter < 0 {
		return errors.New("Program Counter is before end of program")
	}
	if c.ProgramCounter >= len(c.Instructions) {
		return EndOfProgram
	}

	instruction := c.Instructions[c.ProgramCounter]
	// fmt.Printf("Executing \"%v\" at PC %v\n", instruction, c.ProgramCounter)
	switch instruction.Code {
	case "acc":
		c.Accumulator += instruction.Arg
		c.ProgramCounter++
	case "nop":
		c.ProgramCounter++
	case "jmp":
		c.ProgramCounter += instruction.Arg
	default:
		return errors.New("Illegal instruction passed")
	}
	return nil
}

func (c *Computer) Run() error {
	for {
		err := c.ExecuteInstruction()
		if err != nil {
			return err
		}
	}
}

func (tc *TrackedComputer) Run() error {
	for {
		err := tc.ExecuteInstruction()
		if err != nil {
			return err
		}
	}
}

func (tc *TrackedComputer) ExecuteInstruction() error {
	if _, ok := tc.Visited[tc.ProgramCounter]; ok {
		return ErrHitInfiniteLoop
	}
	tc.Visited[tc.ProgramCounter] = struct{}{}
	err := tc.Computer.ExecuteInstruction()
	return err
}

func (c *Computer) Reset() {
	c.Accumulator = 0
	c.ProgramCounter = 0
}

func (tc *TrackedComputer) Reset() {
	tc.Visited = make(map[int]struct{})
	tc.Computer.Reset()
}
