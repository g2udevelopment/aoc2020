package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
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

type Range struct {
	min int
	max int
}

func (r *Range) Update(chardown string, char string, num int) {
	if string(char) == chardown {
		r.max = r.max - num
	} else {
		r.min = r.min + num
	}
}

func parseLine(line string) (int, int) {
	row := Range{min: 0, max: 127}
	col := Range{min: 0, max: 7}
	for idx, char := range line {
		if idx < 7 {
			num := powInt(2, 6-idx)
			row.Update("F", string(char), num)
		} else {
			num := powInt(2, 9-idx)
			col.Update("L", string(char), num)
		}
	}
	return row.max, col.max
}

func CalcId(row int, col int) int {
	return row*8 + col
}
