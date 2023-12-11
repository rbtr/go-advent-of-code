package main

import (
	"fmt"
	"log"
	"strconv"

	common "github.com/rbtr/go-advent-of-code/2021"
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

func stringsToInts(strings ...[]string) ([][]int, error) {
	ints := make([][]int, len(strings))
	var err error
	for i := range strings {
		ints[i] = make([]int, len(strings[i]))
		for j := range strings[i] {
			if ints[i][j], err = strconv.Atoi(strings[i][j]); err != nil {
				return nil, err
			}
		}
	}
	return ints, nil
}

func sumDrops(ints ...[]int) []int {
	res := make([]int, len(ints))
	for i := range ints {
		for j := 1; j < len(ints[i]); j++ {
			if ints[i][j-1] < ints[i][j] {
				res[i]++
			}
		}
	}
	return res
}

func one(puzzle *common.Puzzle) error {
	lines, err := common.SplitLines(puzzle.Sample, puzzle.Input)
	if err != nil {
		return err
	}
	ints, err := stringsToInts(lines...)
	if err != nil {
		return err
	}

	fmt.Printf("solutions: %v\n", sumDrops(ints...))
	return nil
}

func windowedSumDrops(ints ...[]int) []int {
	res := make([]int, len(ints))
	for i := range ints {
		for j := 1; j < len(ints[i])-2; j++ {
			a := ints[i][j-1] + ints[i][j] + ints[i][j+1]
			b := ints[i][j] + ints[i][j+1] + ints[i][j+2]
			if a < b {
				res[i]++
			}
		}
	}
	return res
}

func two(puzzle *common.Puzzle) error {
	lines, err := common.SplitLines(puzzle.Sample, puzzle.Input)
	if err != nil {
		return err
	}
	ints, err := stringsToInts(lines...)
	if err != nil {
		return err
	}

	fmt.Printf("solutions: %v\n", windowedSumDrops(ints...))
	return nil
}
