package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid, _ := readInput("input.txt")

	fmt.Println(grid.doSlope([2]int{3, 1}))
	count := 1

	slopes := []([2]int){{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	for _, slope := range slopes {
		count *= grid.doSlope(slope)
	}

	fmt.Println(count)
}

type grid [][]bool

func (g grid) doSlope(slope [2]int) int {
	count := 0
	x := 0
	for y := 0; y < len(g); y += slope[1] {

		if y == 0 {

		} else {
			x += slope[0]
			if g[y][x%len(g[y])] {
				count++
			}
		}

	}
	return count

}

func readInput(path string) (grid, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var g grid
	for scanner.Scan() {
		var row []bool
		line := scanner.Text()

		for _, char := range line {
			if char == '#' {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		g = append(g, row)

	}
	return g, scanner.Err()
}
