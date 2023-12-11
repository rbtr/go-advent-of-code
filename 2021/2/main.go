package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	common "github.com/rbtr/go-advent-of-code/2021"
)

const (
	forward = "forward"
	down    = "down"
	up      = "up"
)

func main() {
	puzzle, err := common.Load()
	if err != nil {
		log.Fatal(err)
	}
	if err := one(puzzle); err != nil {
		log.Fatal(err)
	}
	if err := two(puzzle); err != nil {
		log.Fatal(err)
	}
}

func calcPositions(ins ...[]string) ([][]int, error) {
	res := make([][]int, len(ins))
	for i := range ins {
		res[i] = []int{0, 0, 0}
		for j := range ins[i] {
			tokens := strings.Split(ins[i][j], " ")
			val, err := strconv.Atoi(tokens[1])
			if err != nil {
				return nil, err
			}
			switch tokens[0] {
			case forward:
				res[i][0] += val
			case down:
				res[i][1] += val
			case up:
				res[i][1] -= val
			}
		}
		res[i][2] = res[i][0] * res[i][1]
	}
	return res, nil
}

func one(puzzle *common.Puzzle) error {
	lines, err := common.SplitLines(puzzle.Sample, puzzle.Input)
	if err != nil {
		return err
	}
	pos, err := calcPositions(lines...)
	if err != nil {
		return err
	}

	fmt.Printf("solutions: %v\n", pos)
	return nil
}

func calcPositionsTraj(ins ...[]string) ([][]int, error) {
	res := make([][]int, len(ins))
	for i := range ins {
		res[i] = []int{0, 0, 0, 0}
		for j := range ins[i] {
			tokens := strings.Split(ins[i][j], " ")
			val, err := strconv.Atoi(tokens[1])
			if err != nil {
				return nil, err
			}
			switch tokens[0] {
			case forward:
				res[i][0] += val
				res[i][1] += val * res[i][2]
			case down:
				res[i][2] += val
			case up:
				res[i][2] -= val
			}
		}
		res[i][3] = res[i][0] * res[i][1]
	}
	return res, nil
}

func two(puzzle *common.Puzzle) error {
	lines, err := common.SplitLines(puzzle.Sample, puzzle.Input)
	if err != nil {
		return err
	}
	pos, err := calcPositionsTraj(lines...)
	if err != nil {
		return err
	}

	fmt.Printf("solutions: %v\n", pos)
	return nil
}
