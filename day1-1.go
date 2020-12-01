package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines, _ := readLines("./input-day1")
	var result int
	for idx1, line1 := range lines {
		for idx2, line2 := range remove(lines, idx1) {
			for _, line3 := range remove(lines, idx2) {
				nline1, _ := strconv.Atoi(line1)
				nline2, _ := strconv.Atoi(line2)
				nline3, _ := strconv.Atoi(line3)
				if nline1+nline2+nline3 == 2020 {
					result = nline1 * nline2 * nline3
					break
				}
			}

		}

	}
	fmt.Printf("The results is: %v", result)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func remove(s []string, index int) []string {
	tmp := make([]string, len(s))
	copy(tmp, s)
	return append(tmp[:index], tmp[index+1:]...)
}
