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

func toBinary(s string) string {
	out := ""
	for _, r := range s {
		switch r {
		case '#':
			out += "1"
		case '.':
			out += "0"
		}
	}
	return out
}

func binStringToDec(s string) (int, error) {
	i, err := strconv.ParseInt(s, 2, 64)
	return int(i), err
}

func countLit(s [][]string) int {
	lit := 0
	for i := range s {
		for j := range s[i] {
			if s[i][j] == "#" {
				lit++
			}
		}
	}
	return lit
}

// map xy coordinates in a larger grid on to a grid smaller by 1 idx
// on all sides.
// inset = 1
func remap(inset, i, j int, in [][]string) (int, int) {
	if i > inset-1 && i <= len(in)+inset-1 {
		i -= inset
	} else {
		i = -1
	}
	if j > inset-1 && j <= len(in[0])+inset-1 {
		j -= inset
	} else {
		j = -1
	}
	return i, j
}

// lookup the prior pixel at x, y and return it
// or the "off" value if it's out-of-bounds in the original grid.
func lookup(iter, i, j int, in [][]string) string {
	ii, jj := remap(1, i, j, in)
	if ii < 0 || jj < 0 {
		if iter%2 == 0 {
			return "#"
		}
		return "."
	}
	return in[ii][jj]
}

func enhance(iter int, algo string, in [][]string) ([][]string, error) {
	inset := 1
	out := make([][]string, len(in)+inset*2)
	for i := range out {
		out[i] = make([]string, len(in[0])+inset*2)
		for j := range out[i] {
			s := ""
			for ii := i - 1; ii <= i+1; ii++ {
				for jj := j - 1; jj <= j+1; jj++ {
					s += lookup(iter, ii, jj, in)
				}
			}
			bin := toBinary(s)
			dec, err := binStringToDec(bin)
			if err != nil {
				return nil, err
			}
			out[i][j] = string(algo[dec])
		}
	}
	return out, nil
}

func one(puzzle common.Data) (string, error) {
	lines := puzzle.AsLines()
	algo := lines[0]
	image, err := common.FromStrings(lines[2:]...).AsStringMatrix()
	if err != nil {
		return "", err
	}
	for i := 0; i < 2; i++ {
		image, err = enhance(i, algo, image)
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("lit: %d", countLit(image)), nil
}

func two(puzzle common.Data) (string, error) {
	lines := puzzle.AsLines()
	algo := lines[0]
	image, err := common.FromStrings(lines[2:]...).AsStringMatrix()
	if err != nil {
		return "", err
	}
	for i := 0; i < 50; i++ {
		image, err = enhance(i, algo, image)
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("lit: %d", countLit(image)), nil
}
