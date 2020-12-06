package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines, _ := readInput("input.txt")
	count := 0

	for _, group := range lines {
		var answerset map[string]struct{} = make(map[string]struct{})
		for _, person := range group {
			for _, answer := range person {
				answerset[string(answer)] = struct{}{}
			}
		}
		count += len(answerset)
	}

	fmt.Println(count)

	count = 0

	for _, group := range lines {
		var answerset map[string]int = make(map[string]int)
		for _, person := range group {
			for _, answer := range person {
				answerset[string(answer)]++
			}
		}
		count2 := 0
		for _, answer := range answerset {
			if answer == len(group) {
				count2++
			}
		}
		count += count2
	}

	fmt.Println(count)

}

func readInput(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(splitOnBlanks)
	var total [][]string
	for scanner.Scan() {
		var answers []string
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		answers = strings.Split(line, "\n")

		total = append(total, answers)

	}
	return total, scanner.Err()
}

func splitOnBlanks(data []byte, atEOF bool) (advance int, token []byte, err error) {

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		return i + 2, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}
