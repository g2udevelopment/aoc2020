package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Policy struct {
	firstOcc  int
	secondOcc int
	letter    string
	password  string
}

// ISValid checks is policy is valid
func (p Policy) IsValid() bool {
	count := strings.Count(p.password, p.letter)
	return count >= p.firstOcc && count <= p.secondOcc
}

func (p Policy) IsValidSecondPolicy() bool {
	valid := false
	for idx, letter := range p.password {
		tdx := idx + 1
		if string(letter) == p.letter {
			if !valid && (tdx == p.firstOcc || tdx == p.secondOcc) {
				valid = true
			} else if valid && (tdx == p.secondOcc) {
				valid = false
				break
			}

		}
	}
	return valid
}

func main() {
	lines, _ := readLines("./input-day2")
	var count1 int
	var count2 int
	for _, line := range lines {
		pol := parseLine(line)
		if pol.IsValid() {
			count1++
		}
		if pol.IsValidSecondPolicy() {
			count2++
		}
	}
	fmt.Printf("Number of valid passwords 1st pol: %v, 2nd pol: %v", count1, count2)
}

func parseLine(line string) Policy {
	split := strings.Split(line, " ")
	boundary := strings.Split(split[0], "-")
	min, _ := strconv.Atoi(boundary[0])
	max, _ := strconv.Atoi(boundary[1])
	letter := split[1][0] //it is ascii so don't bother with runes
	pwd := split[2]
	return Policy{firstOcc: min, secondOcc: max, letter: string(letter), password: pwd}
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
