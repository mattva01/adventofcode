package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/mattva01/adventofcode/2020/8/pkg/compiler"
)

func main() {
	program, _ := readInput("input.txt")
	a := compiler.NewTrackedComputer()
	err := a.LoadProgram(program)
	if err != nil {
		log.Fatal(err)
	}

	err = a.Run()
	if err != nil {
		if err == compiler.ErrHitInfiniteLoop {
			fmt.Println(a.Accumulator)
		} else {
			log.Fatal(err)
		}
	}
	var originalVisited []int
	for k := range a.Visited {
		originalVisited = append(originalVisited, k)
	}
	for done := false; !done; {
		a.Reset()
		err = a.Run()
		if err == compiler.EndOfProgram {
			done = true
		} else {
			if err == compiler.ErrHitInfiniteLoop {
				if len(originalVisited) == 0 {
					log.Fatal("No possible answer")
				}
				n := len(originalVisited) - 1
				linenum := originalVisited[n]
				a.LoadProgram(program)
				permute(a, linenum)
				originalVisited = originalVisited[:n]

			} else {
				log.Fatal(err)
			}
		}

	}
	fmt.Println(a.Accumulator)

}

func permute(tc *compiler.TrackedComputer, linenum int) {
	instruction := &tc.Instructions[linenum]
	switch instruction.Code {
	case "nop":
		instruction.Code = "jmp"
	case "jmp":
		instruction.Code = "nop"
	}

}

func readInput(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

	}
	return lines, scanner.Err()
}
