package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, _ := readInput("input.txt")
	rules := linesToBagRules(lines)
	answer := 0
	for rule, _ := range rules {
		if rules.DeepContain(rule, "shiny gold") {
			answer++
		}
	}
	fmt.Println(answer)

	// Remove 1 because we don't count our initial bag
	fmt.Println(rules.totalBags("shiny gold") - 1)
}

type BagDescription string

type BagRules map[BagDescription]map[BagDescription]int

func (br BagRules) DeepContain(root, input BagDescription) bool {
	if numBags, ok := br[root][input]; ok {
		if numBags > 0 {
			return true
		}
	}

	for x := range br[root] {
		if br.DeepContain(x, input) {
			return true
		}

	}
	return false

}

func (br BagRules) totalBags(bd BagDescription) int {
	bags := 1

	for rule, subBagNum := range br[bd] {
		bags += subBagNum * br.totalBags(rule)
	}

	return bags
}

func linesToBagRules(input []string) BagRules {
	rules := BagRules{}

	for _, rule := range input {
		splitrule := strings.Split(rule, " bags contain ")
		description := BagDescription(splitrule[0])
		contains := make(map[BagDescription]int)
		fmt.Println(splitrule)

		for _, container := range strings.Split(splitrule[1], ",") {
			container = strings.TrimSpace(container)
			if container == "no other bags" {
				continue
			}

			split := strings.Split(container, " ")

			num, _ := strconv.Atoi(split[0])

			contains[BagDescription(strings.Join(split[1:3], " "))] = num

		}
		rules[description] = contains
	}

	return rules

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
