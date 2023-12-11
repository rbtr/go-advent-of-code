package main

import (
	"fmt"
	"log"
	"math"
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

type cube struct {
	x, y, z int
}

type sparse map[cube]interface{}

type span struct {
	start, end int
}

type instr struct {
	on      bool
	x, y, z span
}

func parseInstrs(lines []string) ([]instr, error) {
	instrs := []instr{}
	for _, line := range lines {
		current := instr{}
		tok := strings.Split(line, " ")
		if tok[0] == "on" {
			current.on = true
		}
		spans := strings.Split(tok[1], ",")
		for i := range spans {
			xyzspan := strings.Split(spans[i], "=")
			spanstr := strings.Split(xyzspan[1], "..")
			start, err := strconv.Atoi(spanstr[0])
			if err != nil {
				return nil, err
			}
			end, err := strconv.Atoi(spanstr[1])
			if err != nil {
				return nil, err
			}
			span := span{start: start, end: end}

			switch xyzspan[0] {
			case "x":
				current.x = span
			case "y":
				current.y = span
			case "z":
				current.z = span
			}
		}
		instrs = append(instrs, current)
	}
	return instrs, nil
}

func one(puzzle common.Data) (string, error) {
	lines := puzzle.AsLines()
	instrs, err := parseInstrs(lines)
	if err != nil {
		return "", nil
	}
	ons := sparse{}
	for _, instr := range instrs {
		for x := instr.x.start; x <= instr.x.end; x++ {
			if math.Abs(float64(x)) > 50 {
				continue
			}
			for y := instr.y.start; y <= instr.y.end; y++ {
				if math.Abs(float64(y)) > 50 {
					continue
				}
				for z := instr.z.start; z <= instr.z.end; z++ {
					if math.Abs(float64(z)) > 50 {
						continue
					}
					c := cube{x: x, y: y, z: z}
					// fmt.Printf("%+v -> %+v\n", instr, c)
					if !instr.on {
						delete(ons, c)
					} else {
						ons[c] = nil
					}
				}
			}
		}
	}
	return fmt.Sprintf("on: %d", len(ons)), nil
}

func two(puzzle common.Data) (string, error) {
	lines := puzzle.AsLines()
	instrs, err := parseInstrs(lines)
	if err != nil {
		return "", nil
	}
	ons := sparse{}
	for _, instr := range instrs {
		for x := instr.x.start; x <= instr.x.end; x++ {
			// if math.Abs(float64(x)) > 50 {
			// 	continue
			// }
			for y := instr.y.start; y <= instr.y.end; y++ {
				// if math.Abs(float64(y)) > 50 {
				// 	continue
				// }
				for z := instr.z.start; z <= instr.z.end; z++ {
					// if math.Abs(float64(z)) > 50 {
					// 	continue
					// }
					c := cube{x: x, y: y, z: z}
					// fmt.Printf("%+v -> %+v\n", instr, c)
					if !instr.on {
						delete(ons, c)
					} else {
						ons[c] = nil
					}
				}
			}
		}
	}
	return fmt.Sprintf("on: %d", len(ons)), nil
}
