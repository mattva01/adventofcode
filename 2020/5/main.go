package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	lines, _ := readInput("input.txt")
	var pidList []int
	var maxPid uint = 0
	for _, mask := range lines {
		var rowbyte uint
		var colbyte uint

		for i, maskchar := range mask {
			switch maskchar {
			case 'B':
				rowbyte |= (1 << (i - 3))
			case 'R':
				colbyte |= (1 << i)
			}

		}

		pid := rowbyte*8 + colbyte
		pidList = append(pidList, int(pid))

		if pid > maxPid {
			maxPid = pid
		}
	}
	sort.Ints(pidList)
	var missing int
	for i := pidList[0]; i <= pidList[len(pidList)-1]; i++ {
		test := sort.SearchInts(pidList, i)
		if test < len(pidList) && pidList[test] == i {
		} else {
			missing = i
		}
	}
	fmt.Println(maxPid)
	fmt.Println(missing)

}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
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
		line = Reverse(line)

		lines = append(lines, line)

	}
	return lines, nil
}
