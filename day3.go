package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	EMPTY = iota
	TREE
)

type Instruction struct {
	Right int
	Down  int
}

var prob []Instruction = []Instruction{Instruction{1, 1}, Instruction{3, 1}, Instruction{5, 1}, Instruction{7, 1}, Instruction{1, 2}}

type Toboggan struct {
	Rows       []*Row
	NumTrees   int
	CurrentRow int
	CurrentCol int
	RowL       int
}

func (t *Toboggan) Reset() {
	t.NumTrees = 0
	t.CurrentCol = 0
	t.CurrentRow = 0
	//t.Rows = make([]*Row, 0)
}

type Row struct {
	Objects []int
}

func (t *Toboggan) Drive(r int, d int) {
	t.CurrentRow += d // Set next row
	t.CurrentCol += r
	if t.CurrentRow < t.RowL {
		row := t.Rows[t.CurrentRow]
		l := len(row.Objects)
		v := row.Objects[t.CurrentCol%l]

		if v == TREE {
			t.NumTrees++
		}
	}
}

func NewFromLines(lines []string) *Toboggan {
	t := new(Toboggan)
	for _, line := range lines {
		r := parseRingLine(line)
		t.Rows = append(t.Rows, r)
		t.RowL++
	}

	return t
}

func main() {
	lines, _ := readLines("./input-day3")
	l := len(lines)
	t := NewFromLines(lines)
	var total int = 1
	for _, inst := range prob {
		t.Reset()
		for i := 0; i < l; i++ {
			t.Drive(inst.Right, inst.Down)
		}
		total = total * t.NumTrees
	}
	t.Reset()
	for i := 0; i < l; i++ {
		t.Drive(3, 1)
	}
	fmt.Printf("I encoutered: %v trees", t.NumTrees)
	fmt.Printf("Total trees encountered: %v", total) // Second sol
}

func parseRingLine(line string) *Row {
	//l := len(line)
	r := new(Row)
	for _, char := range line {
		o := parseObject(string(char))
		r.Objects = append(r.Objects, o)
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
