package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	common "github.com/rbtr/go-advent-of-code/2021"
)

func main() {
	puzzle, err := common.Load()
	if err != nil {
		log.Fatal(err)
	}
	common.Solve(puzzle, one, two)
}

func one(p common.Data) (string, error) {
	lines := p.AsLines()
	total := 0
	for i := range lines {
		first, last := "", ""
		for j := range lines[i] {
			if _, err := strconv.Atoi(string(lines[i][j])); err == nil {
				if first == "" {
					first = string(lines[i][j])
				}
				last = string(lines[i][j])
			}
		}
		// fmt.Println(first + last)
		v, _ := strconv.Atoi(first + last)
		total += v
	}
	return fmt.Sprintf("%d", total), nil
}

var digits []string = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func two(p common.Data) (string, error) {
	lines := p.AsLines()

	for i := range lines {
		var line string
		for j := 0; j < len(lines[i]); {
			if _, err := strconv.Atoi(string(lines[i][j])); err == nil {
				line += string(lines[i][j])
				j++
				continue
			}
			for l := range digits {
				if strings.HasPrefix(lines[i][j:], digits[l]) {
					line += fmt.Sprintf("%d", l+1)
				}
			}
			j++
		}
		lines[i] = line
	}
	fmt.Println(lines)

	total := 0
	for i := range lines {
		first, last := "", ""
		for j := range lines[i] {
			if _, err := strconv.Atoi(string(lines[i][j])); err == nil {
				if first == "" {
					first = string(lines[i][j])
				}
				last = string(lines[i][j])
			}
		}
		// fmt.Println(first + last)
		v, _ := strconv.Atoi(first + last)
		total += v
	}
	return fmt.Sprintf("%d", total), nil
}
