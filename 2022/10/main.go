package main

import (
	"fmt"
	"log"

	"github.com/rbtr/go-aoc/pkg/conversion"
	"github.com/rbtr/go-aoc/pkg/puzzle"
)

func main() {
	p, err := puzzle.Load()
	if err != nil {
		log.Fatal(err)
	}
	if err := p.Run(one, two); err != nil {
		log.Fatal(err)
	}
}

func parseInstr(b []byte) (int, int) {
	tok := conversion.SplitWords(b)
	if string(tok[0]) == "noop" {
		return 0, 1
	}
	dx, _ := conversion.ParseSignedInts(tok[1])
	return dx, 2
}

func one(d puzzle.Raw) (string, error) {
	lines := conversion.SplitLines(d)
	x := []int{1}
	for i := range lines {
		dx, dcycle := parseInstr(lines[i])
		x = append(x, x[len(x)-1])
		if dcycle > 1 {
			x = append(x, x[len(x)-1]+dx)
		}
	}
	check := 19
	sum := 0
	for cycle := range x {
		fmt.Println(cycle, x[cycle])
		if cycle == check {
			check += 40
			str := (cycle + 1) * x[cycle]
			sum += str
			fmt.Println(cycle, x[cycle], str)
		}
		if check > 220 {
			break
		}
	}
	return fmt.Sprintf("%d", sum), nil
}

func two(d puzzle.Raw) (string, error) {
	lines := conversion.SplitLines(d)
	x := []int{1}
	for i := range lines {
		dx, dcycle := parseInstr(lines[i])
		// fmt.Println(dx, dcycle)
		x = append(x, x[len(x)-1])
		if dcycle > 1 {
			x = append(x, x[len(x)-1]+dx)
		}
	}
	s := ""
	for i := range x {
		j := i % 40
		if j == 0 {
			s += "\n"
		}
		if x[i]-1 <= j && x[i]+1 >= j {
			s += "#"
		} else {
			s += "."
		}
	}
	return s, nil
}
