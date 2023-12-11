package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"

	common "github.com/rbtr/aoc2021"
)

type XY common.XY

func main() {
	puzzle, err := common.Load()
	if err != nil {
		log.Fatal(err)
	}
	common.Solve(puzzle, one, two)
}

func fold(in map[XY]interface{}, at XY) map[XY]interface{} {
	out := map[XY]interface{}{}
	for xy := range in {
		if at.Y != 0 && xy.Y > at.Y {
			xy.Y = at.Y - (xy.Y - at.Y)
		}
		if at.X != 0 && xy.X > at.X {
			xy.X = at.X - (xy.X - at.X)
		}
		out[xy] = nil
	}
	return out
}

func one(puzzle []byte) (string, error) {
	sc := bufio.NewScanner(bytes.NewReader(puzzle))
	sc.Split(bufio.ScanLines)
	dots := map[XY]interface{}{}
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			break
		}
		xy := strings.Split(line, ",")
		x, err := strconv.Atoi(xy[0])
		if err != nil {
			return "", err
		}
		y, err := strconv.Atoi(xy[1])
		if err != nil {
			return "", err
		}
		dots[XY{X: x, Y: y}] = nil
	}

	instrs := []XY{}
	for sc.Scan() {
		instr := strings.Split(strings.Split(sc.Text(), " ")[2], "=")
		xy := instr[0]
		val, err := strconv.Atoi(instr[1])
		if err != nil {
			return "", err
		}
		if strings.ToLower(xy) == "x" {
			instrs = append(instrs, XY{X: val})
		} else {
			instrs = append(instrs, XY{Y: val})
		}
	}
	out := fold(dots, instrs[0])
	return fmt.Sprintf("%d", len(out)), nil
}

func two(puzzle []byte) (string, error) {
	sc := bufio.NewScanner(bytes.NewReader(puzzle))
	sc.Split(bufio.ScanLines)
	dots := map[XY]interface{}{}
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			break
		}
		xy := strings.Split(line, ",")
		x, err := strconv.Atoi(xy[0])
		if err != nil {
			return "", err
		}
		y, err := strconv.Atoi(xy[1])
		if err != nil {
			return "", err
		}
		dots[XY{X: x, Y: y}] = nil
	}

	instrs := []XY{}
	for sc.Scan() {
		instr := strings.Split(strings.Split(sc.Text(), " ")[2], "=")
		xy := instr[0]
		val, err := strconv.Atoi(instr[1])
		if err != nil {
			return "", err
		}
		if strings.ToLower(xy) == "x" {
			instrs = append(instrs, XY{X: val})
		} else {
			instrs = append(instrs, XY{Y: val})
		}
	}
	out := dots
	for _, instr := range instrs {
		out = fold(out, instr)
	}
	max := XY{}
	for xy := range out {
		if xy.X > max.X {
			max.X = xy.X
		}
		if xy.Y > max.Y {
			max.Y = xy.Y
		}
	}
	s := "\n\n\t"
	for i := 0; i <= max.Y; i++ {
		for j := 0; j <= max.X; j++ {
			xy := XY{Y: i, X: j}
			if _, ok := out[xy]; ok {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n\t"
	}
	return s, nil
}
