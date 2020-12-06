package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./input-day6")
	groups := strings.Split(string(input), "\n\n")

	var total int
	var totalPart2 int
	for _, group := range groups {
		numMembers := len(strings.Split(group, "\n"))
		line := strings.ReplaceAll(group, "\n", "")
		chars := make(map[string]int)
		for _, char := range line {
			chars[string(char)]++
			if chars[string(char)] == numMembers {
				totalPart2++
			}
		}
		total = total + len(chars)
	}
	fmt.Printf("Total unique answers in group: %v\n", total)
	fmt.Printf("Total all answered in group: %v\n", totalPart2)

}
