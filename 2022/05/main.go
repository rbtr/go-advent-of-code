package main

import (
	"log"
	"strconv"
	"strings"

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

func parse(d puzzle.Raw) ([][]string, [][]int) {
	lines := conversion.SplitLines(d)
	labelRow := 0
	for labelRow = range lines {
		str, _ := conversion.ParseString(lines[labelRow])
		if strings.TrimSpace(str) == "" {
			labelRow--
			break
		}
	}
	stacks := [][]string{}
	for stack := 1; stack < len(lines[labelRow]); stack += 4 {
		stacks = append(stacks, []string{})
	}
	for row := labelRow - 1; row >= 0; row-- {
		for stack := range stacks {
			entry := string(lines[row][stack*4+1])
			if strings.TrimSpace(entry) == "" {
				continue
			}
			stacks[stack] = append(stacks[stack], entry)
		}
	}
	instructions := [][]int{}
	for i := labelRow + 2; i < len(lines); i++ {
		tokens := strings.Split(string(lines[i]), " ")
		qty, _ := strconv.Atoi(tokens[1])
		from, _ := strconv.Atoi(tokens[3])
		to, _ := strconv.Atoi(tokens[5])
		instructions = append(instructions, []int{qty, from, to})
	}
	return stacks, instructions
}

func one(d puzzle.Raw) (string, error) {
	stacks, instrs := parse(d)
	for _, instr := range instrs {
		for qty := 0; qty < instr[0]; qty++ {
			stacks[instr[2]-1] = append(stacks[instr[2]-1], stacks[instr[1]-1][len(stacks[instr[1]-1])-1])
			stacks[instr[1]-1] = stacks[instr[1]-1][:len(stacks[instr[1]-1])-1]
		}
	}
	tops := ""
	for i := range stacks {
		tops += stacks[i][len(stacks[i])-1]
	}
	return tops, nil
}

func two(d puzzle.Raw) (string, error) {
	stacks, instrs := parse(d)
	for _, instr := range instrs {
		stacks[instr[2]-1] = append(stacks[instr[2]-1], stacks[instr[1]-1][len(stacks[instr[1]-1])-instr[0]:]...)
		stacks[instr[1]-1] = stacks[instr[1]-1][:len(stacks[instr[1]-1])-instr[0]]
	}
	tops := ""
	for i := range stacks {
		tops += stacks[i][len(stacks[i])-1]
	}
	return tops, nil
}
