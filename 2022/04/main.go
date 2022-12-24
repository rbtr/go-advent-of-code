package main

import (
	"fmt"
	"strconv"
	"strings"

	common "github.com/rbtr/aoc2021"
)

func main() {
	one()
	two()
}

func parseTuple(in string) []int {
	indices := strings.Split(in, "-")
	start, _ := strconv.Atoi(indices[0])
	end, _ := strconv.Atoi(indices[1])
	return []int{start, end}
}

func one() {
	p, _ := common.Load()
	lines := p.Sample.AsLines()
	count := 0
	for _, line := range lines {
		rangeStrings := strings.Split(line, ",")
		first := parseTuple(rangeStrings[0])
		second := parseTuple(rangeStrings[1])
		if encompasses(first, second) {
			count++
		}
	}
	fmt.Println(count)
}

func encompasses(first, second []int) bool {
	if first[1]-first[0] < second[1]-second[0] { // put the smaller range second
		first, second = second, first
	}
	return first[0] <= second[0] && first[1] >= second[1]
}

func two() {
	p, _ := common.Load()
	lines := p.Input.AsLines()
	count := 0
	for _, line := range lines {
		rangeStrings := strings.Split(line, ",")
		first := parseTuple(rangeStrings[0])
		second := parseTuple(rangeStrings[1])
		if overlaps(first, second) {
			fmt.Println(line)
			count++
		}
	}
	fmt.Println(count)
}

func overlaps(first, second []int) bool {
	return first[0] >= second[0] && first[0] <= second[1] ||
		first[1] >= second[0] && first[1] <= second[1] ||
		second[0] >= first[0] && second[0] <= first[1]
}
