package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Console struct {
	LinesVisited   map[int]bool
	ProgramCounter int
	Acc            int
	Program        []string
	Halted         bool
	Looped         bool
}

func NewConsole(program []string) Console {
	return Console{make(map[int]bool), 0, 0, program, false, false}
}

// bool -> looped
func (c *Console) Execute() bool {

	for {
		if c.Halted {
			break
		}
		c.NextInstruction()
	}
	return c.Looped

}

func (c *Console) NextInstruction() {
	nextLine := c.Program[c.ProgramCounter]
	c.LinesVisited[c.ProgramCounter] = true
	args := strings.Split(nextLine, " ")[1:]
	if strings.HasPrefix(nextLine, "jmp") {
		c.ProgramCounter += atoi(args[0])
	} else if strings.HasPrefix(nextLine, "acc") {
		c.Acc += atoi(args[0])
		c.ProgramCounter++
	} else if strings.HasPrefix(nextLine, "nop") {
		c.ProgramCounter++
	}
	//end?
	if c.ProgramCounter >= len(c.Program) {
		c.Halted = true
		return
	}
	//looped?
	if c.LinesVisited[c.ProgramCounter] {
		c.Halted = true
		c.Looped = true
	}
}

func main() {
	input, _ := ioutil.ReadFile("./input-day8")
	lines := strings.Split(string(input), "\n")
	p := NewConsole(lines)
	p.Execute()
	fmt.Printf("Program looped: %v, acc: %v\n", p.Looped, p.Acc)

	//part 2
	for idx, line := range lines {
		var clines []string = make([]string, len(lines))
		copy(clines, lines)
		if strings.HasPrefix(line, "jmp") {
			clines[idx] = strings.Replace(line, "jmp", "nop", 1)
		}
		if strings.HasPrefix(line, "nop") {
			clines[idx] = strings.Replace(line, "nop", "jmp", 1)
		}
		//fmt.Println(clines)
		p := NewConsole(clines)
		p.Execute()
		if p.Halted && !p.Looped {
			fmt.Printf("Found correct program mut on line: %v, acc: %v\n", idx, p.Acc)
		}
	}
}

func atoi(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}
