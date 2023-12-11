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
	common.Solve(puzzle, one, two)
}

func anyWin(win int, scores ...int) bool {
	for i := range scores {
		if scores[i] >= win {
			return true
		}
	}
	return false
}

func roll(d, turn int) int {
	out := 0
	for i := 1; i < 4; i++ {
		roll := (3*turn + i) % d
		if roll == 0 {
			roll = d
		}
		out += roll
	}
	return out
}

func move(wrap, current, move int) int {
	next := (current + move) % wrap
	if next == 0 {
		next = wrap
	}
	return next
}

func minimum(ints ...int) int {
	min := ints[0]
	for i := range ints {
		if min > ints[i] {
			min = ints[i]
		}
	}
	return min
}

func one(puzzle common.Data) (string, error) {
	lines := puzzle.AsLines()
	positions := make([]int, len(lines))
	scores := make([]int, len(positions))
	for i := range lines {
		var err error
		if positions[i], err = strconv.Atoi(string(lines[i][len(lines[i])-1])); err != nil {
			return "", err
		}
		scores[i] = 0
	}

	i := 0
	for ; !anyWin(1000, scores...); i++ {
		res := roll(100, i)
		positions[i%2] = move(10, positions[i%2], res)
		scores[i%2] += positions[i%2]
	}
	loser := minimum(scores...)
	return fmt.Sprintf("rolls: %d, loser: %d = %d", i*3, loser, i*3*loser), nil
}

func two(puzzle common.Data) (string, error) {
	lines := puzzle.AsLines()
	positions := make([]int, len(lines))
	scores := make([]int, len(positions))
	for i := range lines {
		var err error
		if positions[i], err = strconv.Atoi(string(lines[i][len(lines[i])-1])); err != nil {
			return "", err
		}
		scores[i] = 0
	}
	i := 0
	for ; !anyWin(21, scores...); i++ {
		res := roll(3, i)
		positions[i%2] = move(10, positions[i%2], res)
		scores[i%2] += positions[i%2]
	}
	loser := minimum(scores...)
	return fmt.Sprintf("rolls: %d, loser: %d = %d", i*3, loser, i*3*loser), nil
}
