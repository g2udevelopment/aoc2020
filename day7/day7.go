package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	BagsCount int    = 12
	BagCount  int    = 8
	NoBags    int    = 7
	GoldBag   string = "shinygold"
)

type Bag struct {
	Count  int
	Color  ColorKey
	Bags   []*Bag
	CurMul int
}

type ColorKey string

var bags map[ColorKey]*Bag = make(map[ColorKey]*Bag)

func main() {
	input, _ := ioutil.ReadFile("./input-day7")
	lines := strings.Split(string(input), "\n")
	var total int
	for _, line := range lines {
		bag := createBag(line)
		bags[bag.Color] = bag
	}

	for _, bag := range bags {
		var q []*Bag = make([]*Bag, 0)
		q = append(q, bag)
		if searchBag(GoldBag, q) {
			total++
		}
	}
	//j, _ := json.Marshal(bags)
	//fmt.Println(string(j))
	var q []*Bag = make([]*Bag, 0)
	q = append(q, bags[ColorKey(GoldBag)])
	count := countBag(q, 0)
	fmt.Printf("Num of goldbags inside: %v\n", total)
	fmt.Printf("Num of bags inside goldbag: %v\n", count)
}

func countBag(q []*Bag, sum int) int {
	if len(q) == 0 {
		return sum
	}
	cur := q[0]
	var bagsum int
	for _, bag := range cur.Bags {
		bagsum += bag.Count
		nbag := bags[bag.Color]
		b := Bag{Count: bag.Count, Bags: nbag.Bags, CurMul: bag.Count * cur.CurMul, Color: nbag.Color}
		q = append(q, &b) //append the next bag as child
	}
	sum += bagsum * cur.CurMul
	return countBag(q[1:], sum)
}

// BFS
func searchBag(k string, q []*Bag) bool {
	if len(q) == 0 {
		return false
	}
	cur := q[0]

	for _, bag := range cur.Bags {
		if bag.Color == ColorKey(k) {
			return true
		}
		q = append(q, bags[bag.Color]) //append the next bag as child
	}

	return searchBag(k, q[1:])
}

func atoi(s string) int {
	value, _ := strconv.Atoi(s)
	return value
}

func createBag(line string) *Bag {
	words := strings.Split(line, " ")
	bag := Bag{Count: 1, Color: ColorKey(words[0] + words[1]), CurMul: 1}
	parts := strings.Split(line, ", ")
	if len(words) > NoBags {
		bag.Bags = append(bag.Bags, &Bag{Color: ColorKey(words[5] + words[6]), Count: atoi(words[4])})
	}
	for _, part := range parts[1:] {
		w := strings.Split(part, " ")
		bag.Bags = append(bag.Bags, &Bag{Color: ColorKey(w[1] + w[2]), Count: atoi(w[0])})
	}
	return &bag
}
