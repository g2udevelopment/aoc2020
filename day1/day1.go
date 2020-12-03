package main

import (
	"aoc2020/util"
	"fmt"
	"strconv"
)

func main() {
	lines, _ := util.ReadLines("./input-day1")
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

func remove(s []string, index int) []string {
	tmp := make([]string, len(s))
	copy(tmp, s)
	return append(tmp[:index], tmp[index+1:]...)
}
