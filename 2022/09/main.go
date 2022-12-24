package main

import (
	"fmt"
	"log"
	"math"

	"github.com/rbtr/go-aoc/pkg/conversion"
	"github.com/rbtr/go-aoc/pkg/puzzle"
	"github.com/rbtr/go-aoc/pkg/series"
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

type move struct {
	direction string
	speed     int
}

type coord struct {
	x, y int
}

func (c coord) Comparable() coord {
	return c
}

func calculateMove(a, b coord) coord {
	dx := a.x - b.x                     // => -2
	dy := a.y - b.y                     // => 0
	absdx := int(math.Abs(float64(dx))) // => 2
	absdy := int(math.Abs(float64(dy))) // => 0
	if absdx <= 1 && absdy <= 1 {
		return b
	}
	if absdx != 0 {
		b.x += dx / absdx
	}
	if absdy != 0 {
		b.y += dy / absdy
	}
	return b
}

func one(d puzzle.Raw) (string, error) {
	moves := []move{}
	lines := conversion.SplitLines(d)
	for i := range lines {
		tokens := conversion.SplitWords(lines[i])
		speed, _ := conversion.ParseInt(tokens[1])
		moves = append(moves, move{direction: string(tokens[0]), speed: speed})
	}
	visited := series.Set[coord, coord]{}
	knots := make([]coord, 2)
	visited.Add(knots[1])
	for _, move := range moves {
		for i := 0; i < move.speed; i++ {
			switch move.direction {
			case "U":
				knots[0].y++
			case "D":
				knots[0].y--
			case "R":
				knots[0].x++
			case "L":
				knots[0].x--
			}
			for j := 0; j < len(knots)-1; j++ {
				knots[j+1] = calculateMove(knots[j], knots[j+1])
			}
			visited.Add(knots[len(knots)-1])
		}
	}
	fmt.Println(visited)
	return fmt.Sprintf("visited: %d", len(visited)), nil
}

func two(d puzzle.Raw) (string, error) {
	moves := []move{}
	lines := conversion.SplitLines(d)
	for i := range lines {
		tokens := conversion.SplitWords(lines[i])
		speed, _ := conversion.ParseInt(tokens[1])
		moves = append(moves, move{direction: string(tokens[0]), speed: speed})
	}
	visited := series.Set[coord, coord]{}

	knots := make([]coord, 10)
	visited.Add(knots[len(knots)-1])
	for _, move := range moves {
		for i := 0; i < move.speed; i++ {
			switch move.direction {
			case "U":
				knots[0].y++
			case "D":
				knots[0].y--
			case "R":
				knots[0].x++
			case "L":
				knots[0].x--
			}
			for j := 0; j < len(knots)-1; j++ {
				knots[j+1] = calculateMove(knots[j], knots[j+1])
			}
			visited.Add(knots[len(knots)-1])
		}
	}
	return fmt.Sprintf("visited: %d", len(visited)), nil
}

func draw(visited series.Set[coord, coord]) {
	maxx, minx, maxy, miny := 0, 0, 0, 0
	for visit := range visited {
		if visit.x > maxx {
			maxx = visit.x
		}
		if visit.x < minx {
			minx = visit.x
		}
		if visit.y > maxy {
			maxy = visit.y
		}
		if visit.y < miny {
			miny = visit.y
		}
	}
	s := "\n"
	for y := miny; y < maxy+1; y++ {
		line := ""
		for x := minx; x < maxx+1; x++ {
			if visited.Contains(coord{x, y}) {
				line += "#"
			} else {
				line += "."
			}
		}
		s = line + "\n" + s
	}
	s += "\n"
	fmt.Println(s)
}
