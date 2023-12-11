package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	common "github.com/rbtr/aoc2021"
)

func main() {
	puzzle, err := common.Load()
	if err != nil {
		log.Fatal(err)
	}
	common.Solve(puzzle, one, two)
}

type Area struct {
	x0, x1, y0, y1 int
}

func parseTarget(s string) (Area, error) {
	ranges := strings.Split(s, ":")[1]
	splitRanges := strings.Split(ranges, ",")
	xrange := strings.Split(splitRanges[0], "=")
	x0x1 := strings.Split(xrange[1], "..")
	yrange := strings.Split(splitRanges[1], "=")
	y0y1 := strings.Split(yrange[1], "..")
	x0, err := strconv.Atoi(x0x1[0])
	if err != nil {
		return Area{}, err
	}
	x1, err := strconv.Atoi(x0x1[1])
	if err != nil {
		return Area{}, err
	}
	y0, err := strconv.Atoi(y0y1[0])
	if err != nil {
		return Area{}, err
	}
	y1, err := strconv.Atoi(y0y1[1])
	if err != nil {
		return Area{}, err
	}
	return Area{x0: x0, x1: x1, y0: y0, y1: y1}, nil
}

func vymax(target Area) int {
	return int(math.Abs(float64(target.y0))) - 1
}

func triangle(n int) int {
	return n * (n + 1) / 2
}

func one(puzzle common.Data) (string, error) {
	lines := puzzle.SplitLines()
	target, err := parseTarget(lines[0])
	if err != nil {
		return "", err
	}
	vy := vymax(target)
	return fmt.Sprintf("vy: %d, h: %d", vy, triangle(vy)), nil
}

func vymin(target Area) int {
	return int(math.Abs(float64(target.y1))) - 1
}

func t(v int) int {
	return v/2 + 1
}

func vx(x, t int) int {
	return (x + (t*t)/2) / t
}

func vxs(target Area, t int) []int {
	out := map[int]interface{}{}
	for x := target.x0; x < target.x1; x++ {
		v := vx(x, t)
		out[v] = nil
	}
	ints := []int{}
	for k := range out {
		ints = append(ints, k)
	}
	return ints
}

type XY common.XY

func two(puzzle common.Data) (string, error) {
	lines := puzzle.SplitLines()
	target, err := parseTarget(lines[0])
	if err != nil {
		return "", err
	}
	vymax := vymax(target)
	vymin := vymin(target)

	trajs := map[XY]interface{}{}
	for vy := vymax; vy > vymin; vy-- {
		t := t(vy)
		vxs := vxs(target, t)
		for _, vx := range vxs {
			trajs[XY{X: vx, Y: vy}] = nil
		}
	}

	for vy := target.y0; vy > target.y1; vy-- {
		t := t()
	}
	// fmt.Printf("%v", trajs)

	return fmt.Sprintf("num: %d", len(trajs)), nil
}
