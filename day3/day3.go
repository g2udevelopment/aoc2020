package main

import (
	"aoc2020/util"
	"fmt"
)

const (
	EMPTY = iota
	TREE
)

type Instruction struct {
	Right int
	Down  int
}

var prob []Instruction = []Instruction{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

type Toboggan struct {
	Map *util.Map
}

func (t *Toboggan) Drive(r int, d int, symbol int) int {
	var x, y, count int
	for i := 0; i < t.Map.Height; i++ {
		x += r
		y += d
		if y < t.Map.Height && t.Map.SymbolAtWrappedRight(x, y) == symbol {

			count++
		}
	}
	return count
}

func NewFromLines(lines []string) *Toboggan {
	t := new(Toboggan)
	t.Map = new(util.Map)
	for _, line := range lines {
		r := parseRingLine(line)
		t.Map.Rows = append(t.Map.Rows, r)
	}
	t.Map.Height = len(lines)
	t.Map.Width = len(t.Map.Rows[0])
	return t
}

func main() {
	lines, _ := util.ReadLines("./input-day3")
	t := NewFromLines(lines)
	var second int = 1

	first := t.Drive(3, 1, TREE)

	for _, inst := range prob {
		count := t.Drive(inst.Right, inst.Down, TREE)
		second = second * count
	}

	fmt.Printf("I encoutered: %v trees \n", first)
	fmt.Printf("Total trees encountered: %v", second) // Second sol
}

func parseRingLine(line string) []int {
	r := make([]int, 0)
	for _, char := range line {
		o := parseObject(string(char))
		r = append(r, o)
	}
	return r
}

func parseObject(o string) int {
	switch o {
	case ".":
		return EMPTY
	case "#":
		return TREE
	default:
		return -1
	}
}
