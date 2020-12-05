package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("./input-day5")
	lines := strings.Split(string(input), "\n")
	fmt.Printf("num lines: %v\n", len(lines))
	var maxID int
	var ids []int
	for _, line := range lines {
		//fmt.Println(CalcId(parseLine(line)))
		id := CalcId(parseLine(line))
		ids = append(ids, id)
		if id > maxID {
			maxID = id
		}
	}

	sort.Ints(ids)
	//fmt.Println(ids)
	for idx, num := range ids {
		if idx != len(ids)-1 {
			if ids[idx+1]-num > 1 {
				fmt.Printf("My id is %v: ", num+1)
			}
		}
	}

	fmt.Printf("Highest id: %v", maxID)
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func parseLine(line string) (int, int) {
	row, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(line[0:7], "F", "0"), "B", "1"), 2, 32)
	col, _ := strconv.ParseInt(strings.ReplaceAll(strings.ReplaceAll(line[7:10], "L", "0"), "R", "1"), 2, 32)

	return int(row), int(col)
}

func CalcId(row int, col int) int {
	return row*8 + col
}
