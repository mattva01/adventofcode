package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	count := 0
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, ppp := range input {
		if validate(ppp) {
			count++
		}
	}
	fmt.Println(count)
	count = 0
	for _, ppp := range input {
		if validate2(ppp) {
			count++
		}
	}
	fmt.Println(count)
}

type PasswordPolicy struct {
	Character string
	MinTimes  int
	MaxTimes  int
}

type PolicyPasswordPair struct {
	Policy   PasswordPolicy
	Password string
}

func validate2(input PolicyPasswordPair) bool {
	count := 0
	// length := len(input.Password)
	fmt.Println(input)
	if input.Password[input.Policy.MinTimes-1] == input.Policy.Character[0] {
		count++
	}

	if input.Password[input.Policy.MaxTimes-1] == input.Policy.Character[0] {
		count++
	}
	if count == 1 {
		return true
	}
	return false
}

func validate(input PolicyPasswordPair) bool {
	count := strings.Count(input.Password, input.Policy.Character)
	if count < input.Policy.MinTimes || count > input.Policy.MaxTimes {
		fmt.Printf("%v appears %v times in %v, violating min %v or max %v\n", input.Policy.Character, count, input.Password, input.Policy.MinTimes, input.Policy.MaxTimes)
		return false
	}
	return true
}
func readInput(path string) ([]PolicyPasswordPair, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var output []PolicyPasswordPair
	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile(`^(\d+)-(\d+) (.*): (.*)$`)
	for scanner.Scan() {
		var ppp PolicyPasswordPair
		parts := r.FindStringSubmatch(scanner.Text())
		ppp.Password = parts[4]
		min, _ := strconv.Atoi(parts[1])
		max, _ := strconv.Atoi(parts[2])

		ppp.Policy = PasswordPolicy{
			Character: parts[3],
			MinTimes:  min,
			MaxTimes:  max,
		}
		if err != nil {
			return nil, err
		}
		output = append(output, ppp)

	}
	return output, scanner.Err()
}
