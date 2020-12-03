package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	pair, err := findNumbers(input, 2020)
	if err != nil {
		log.Fatal(err)
	}
	pair2, err := findNumbers2(input, 2020)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(product(pair[:]))
	fmt.Println(product(pair2[:]))

}

func product(slice []int) int {
	result := slice[0]
	for _, value := range slice[1:] {
		result *= value
	}
	return result
}

func findNumbers(problem []int, targetValue int) ([2]int, error) {
	set := make(map[int]struct{})

	for _, value := range problem {
		candidate := targetValue - value
		if _, ok := set[candidate]; ok {
			return [2]int{value, candidate}, nil
		}
		set[value] = struct{}{}
	}
	return [2]int{}, errors.New("Cannot find pair")
}

func findNumbers2(problem []int, targetValue int) ([3]int, error) {

	for i, value := range problem {
		set := make(map[int]struct{})
		initialSum := targetValue - value
		for _, value2 := range problem[i+1:] {
			finalSum := initialSum - value2
			if _, ok := set[finalSum]; ok {
				return [3]int{value, value2, finalSum}, nil
			}
			set[value2] = struct{}{}
		}

	}
	return [3]int{}, errors.New("Cannot find triplet")
}

func readInput(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var output []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		output = append(output, number)

	}
	return output, scanner.Err()
}
