package main

import (
	"aoc2020/util"
	"fmt"
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
	first := string(p.password[p.firstOcc-1]) == p.letter && string(p.password[p.secondOcc-1]) != p.letter
	second := string(p.password[p.firstOcc-1]) != p.letter && string(p.password[p.secondOcc-1]) == p.letter
	return first || second
}

func main() {
	lines, _ := util.ReadLines("./input-day2")
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
