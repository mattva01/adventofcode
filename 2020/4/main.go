package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	rows, _ := readInput("input.txt")
	reqKeys := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	passportreqs := []PassportReq{
		{Key: "byr", Regex: "^(19[2-8][0-9]|199[0-9]|200[0-2])$"},
		{Key: "iyr", Regex: "^(201[0-9]|2020)$"},
		{Key: "eyr", Regex: "^(202[0-9]|2030)$"},
		{Key: "hgt", Regex: "^((1[5-8][0-9]|19[0-3])cm|(59|6[0-9]|7[0-6])in)$"},
		{Key: "hcl", Regex: "^#[0-9a-f]{6}$"},
		{Key: "ecl", Regex: "^(amb|blu|brn|gry|grn|hzl|oth)$"},
		{Key: "pid", Regex: "^[0-9]{9}$"},
	}
	count := 0

	for _, row := range rows {
		valid := true
		for _, key := range reqKeys {
			if _, ok := row[key]; !ok {
				valid = false
			}
		}
		if valid {
			count++
		}
	}
	fmt.Println(count)

	count = 0

	for _, row := range rows {
		valid := true
		for _, req := range passportreqs {
			if _, ok := row[req.Key]; !ok {
				valid = false
			} else {
				match, _ := regexp.MatchString(req.Regex, row[req.Key])
				if !match {
					valid = false
				}
			}
		}
		if valid {
			count++
		}
	}
	fmt.Println(count)

}

type PassportReq struct {
	Key   string
	Regex string
}

type passport map[string]string

func readInput(path string) ([]passport, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(splitOnBlanks)
	var passports []passport
	for scanner.Scan() {
		var pass passport
		pass = make(map[string]string)
		line := scanner.Text()
		line = strings.Replace(line, "\n", " ", -1)

		for _, thing := range strings.Split(line, " ") {
			result := strings.Split(thing, ":")
			pass[result[0]] = result[1]
		}
		passports = append(passports, pass)

	}
	return passports, scanner.Err()
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
