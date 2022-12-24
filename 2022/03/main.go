package main

import (
	"fmt"
	"strings"

	common "github.com/rbtr/aoc2021"
)

var scores = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	one()
	two()
}

func one() {
	p, _ := common.Load()
	lines := p.Input.AsLines()
	sum := 0
	for _, line := range lines {
		lenline := len(line)
		sum += searchMatches(lenline, line)
	}
	fmt.Println(sum)
}

func searchMatches(lenline int, line string) int {
	for i := 0; i < lenline/2; i++ {
		itemOne := line[i]
		for j := lenline / 2; j < lenline; j++ {
			itemTwo := line[j]
			if string(itemOne) == string(itemTwo) {
				return strings.Index(scores, string(itemOne)) + 1
			}
		}
	}
	return 0
}

func two() {
	p, _ := common.Load()
	lines := p.Input.AsLines()
	sums := 0
	for i := 0; i < len(lines); i += 3 {
		first := lines[i]
		second := lines[i+1]
		third := lines[i+2]
		sums += search(first, second, third)
	}
	fmt.Println(sums)
}

func search(first, second, third string) int {
	for j := range first {
		for k := range second {
			if first[j] == second[k] {
				for l := range third {
					if second[k] == third[l] {
						return strings.Index(scores, string(first[j])) + 1
					}
				}
			}
		}
	}
	return 0
}
