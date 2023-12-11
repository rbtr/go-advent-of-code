package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"sort"
	"strconv"

	common "github.com/rbtr/aoc2021"
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

func minimas(b []byte) ([][]int, [][]int, error) {
	sc := bufio.NewScanner(bytes.NewReader(b))
	sc.Split(bufio.ScanLines)
	lines := [][]int{}
	for sc.Scan() {
		sw := bufio.NewScanner(bytes.NewReader(sc.Bytes()))
		sw.Split(bufio.ScanRunes)
		line := []int{}
		for sw.Scan() {
			i, err := strconv.Atoi(sw.Text())
			if err != nil {
				return nil, nil, err
			}
			line = append(line, i)
		}
		lines = append(lines, line)
	}
	minimas := [][]int{}
	for i := range lines {
		for j := range lines[i] {
			val := lines[i][j]
			// up
			if i > 0 {
				up := lines[i-1][j]
				if val >= up {
					continue
				}
			}
			// down
			if i < len(lines)-1 {
				down := lines[i+1][j]
				if val >= down {
					continue
				}
			}
			// left
			if j > 0 {
				left := lines[i][j-1]
				if val >= left {
					continue
				}
			}
			// right
			if j < len(lines[i])-1 {
				right := lines[i][j+1]
				if val >= right {
					continue
				}
			}
			minimas = append(minimas, []int{i, j})
		}
	}
	return lines, minimas, nil
}

func one(puzzle *common.Puzzle) error {
	lines, minimas, err := minimas(puzzle.Input)
	if err != nil {
		return err
	}
	sum := 0
	for _, xy := range minimas {
		sum = sum + 1 + lines[xy[0]][xy[1]]
	}
	fmt.Printf("height: %d\n", sum)
	return nil
}

func deepCopy(in [][]int) [][]int {
	out := make([][]int, len(in))
	for i := range in {
		out[i] = make([]int, len(in[i]))
		for j := range in[i] {
			out[i][j] = in[i][j]
		}
	}
	return out
}

func findBasin(lines [][]int, xy []int) [][]int {
	out := [][]int{xy}
	lines[xy[0]][xy[1]] = 9
	// up
	if xy[0] > 0 {
		up := lines[xy[0]-1][xy[1]]
		if up < 9 {
			out = append(out, findBasin(lines, []int{xy[0] - 1, xy[1]})...)
		}
	}
	// left
	if xy[1] > 0 {
		left := lines[xy[0]][xy[1]-1]
		if left < 9 {
			out = append(out, findBasin(lines, []int{xy[0], xy[1] - 1})...)
		}
	}
	// down
	if xy[0] < len(lines)-1 {
		down := lines[xy[0]+1][xy[1]]
		if down < 9 {
			out = append(out, findBasin(lines, []int{xy[0] + 1, xy[1]})...)
		}
	}
	// right
	if xy[1] < len(lines[xy[0]])-1 {
		right := lines[xy[0]][xy[1]+1]
		if right < 9 {
			out = append(out, findBasin(lines, []int{xy[0], xy[1] + 1})...)
		}
	}
	return out
}

func two(puzzle *common.Puzzle) error {
	lines, minimas, err := minimas(puzzle.Input)
	if err != nil {
		return err
	}
	basins := [][][]int{}
	for _, xy := range minimas {
		basins = append(basins, findBasin(lines, xy))
	}
	lens := []int{}
	for i := range basins {
		lens = append(lens, len(basins[i]))
	}
	sort.Ints(lens)
	fmt.Println(lens[len(lens)-1] * lens[len(lens)-2] * lens[len(lens)-3])
	return nil
}
