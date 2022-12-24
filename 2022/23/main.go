package main

import (
	"fmt"
	"log"
	"math"

	"github.com/rbtr/go-aoc/pkg/conversion"
	"github.com/rbtr/go-aoc/pkg/puzzle"
	"golang.org/x/exp/maps"
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

type coordinate struct {
	x, y int
}

type elf struct {
	loc coordinate
}

func ingest(d puzzle.Raw) (map[coordinate]*elf, coordinate) {
	elves := map[coordinate]*elf{}
	extents := coordinate{}
	lines := conversion.SplitLines(d)
	for y, line := range lines {
		chars := conversion.SplitCharacters(line)
		for x, char := range chars {
			if string(char) == "#" {
				loc := coordinate{x: x, y: y}
				elves[loc] = &elf{loc}
				if x > extents.x {
					extents.x = x
				}
				if y > extents.y {
					extents.y = y
				}
			}
		}
	}
	return elves, extents
}

func adjOccupiedFunc(direction coordinate) func(map[coordinate]*elf, coordinate) (coordinate, bool) {
	sweep := []coordinate{direction, direction, direction}
	if direction.x == 0 {
		sweep[0].x, sweep[2].x = -1, 1
	}
	if direction.y == 0 {
		sweep[0].y, sweep[2].y = -1, 1
	}
	return func(elves map[coordinate]*elf, current coordinate) (coordinate, bool) {
		for _, delta := range sweep {
			delta.x += current.x
			delta.y += current.y
			// fmt.Println("delta", delta)
			if _, ok := elves[delta]; ok {
				return coordinate{}, true
			}
		}
		return direction, false
	}
}

func findRange(coords []coordinate) []coordinate {
	extents := []coordinate{{x: math.MaxInt, y: math.MaxInt}, {x: math.MinInt, y: math.MinInt}}
	for _, coord := range coords {
		if coord.x < extents[0].x {
			extents[0].x = coord.x
		}
		if coord.x > extents[1].x {
			extents[1].x = coord.x
		}
		if coord.y < extents[0].y {
			extents[0].y = coord.y
		}
		if coord.y > extents[1].y {
			extents[1].y = coord.y
		}
	}
	return extents
}

func draw(elves map[coordinate]*elf) {
	extents := findRange(maps.Keys(elves))
	s := "\n"
	for y := extents[0].y; y <= extents[1].y; y++ {
		for x := extents[0].x; x <= extents[1].x; x++ {
			if _, ok := elves[coordinate{x: x, y: y}]; ok {
				s += "#"
				continue
			}
			s += "."
		}
		s += "\n"
	}
	fmt.Println(s)
}

func one(d puzzle.Raw) (string, error) {
	elves, _ := ingest(d)
	nswe := [](func(map[coordinate]*elf, coordinate) (coordinate, bool)){
		adjOccupiedFunc(coordinate{x: 0, y: -1}),
		adjOccupiedFunc(coordinate{x: 0, y: 1}),
		adjOccupiedFunc(coordinate{x: -1, y: 0}),
		adjOccupiedFunc(coordinate{x: 1, y: 0}),
	}
	turn := 0
	for {
		if turn == 10 {
			break
		}
		moves := map[coordinate][]*elf{}
		for current, elf := range elves {
			// fmt.Println(current)
			var next *coordinate
			unblocked := true
			for i := range nswe {
				direction, blocked := nswe[(turn+i)%4](elves, current)
				if blocked {
					unblocked = false
					// fmt.Println("unblocked", unblocked)
					continue
				}
				if next == nil {
					// fmt.Println("first")
					next = &direction
				}
			}
			if unblocked {
				// fmt.Println("unblocked continue")
				continue
			}
			if next == nil {
				// fmt.Println("next nil")
				continue
			}
			current.x += next.x
			current.y += next.y
			moves[current] = append(moves[current], elf)
		}
		// fmt.Println(turn, moves)
		for move, candidates := range moves {
			if len(candidates) > 1 {
				continue
			}
			elf := candidates[0]
			delete(elves, elf.loc)
			elf.loc = move
			elves[move] = elf
		}
		draw(elves)
		turn++
	}

	extent := findRange(maps.Keys(elves))
	x := int(math.Abs(float64(extent[1].x - extent[0].x)))
	y := int(math.Abs(float64(extent[1].y - extent[0].y)))
	return fmt.Sprintf("%v", ((x+1)*(y+1))-len(elves)), nil
}

func two(d puzzle.Raw) (string, error) {
	elves, _ := ingest(d)
	nswe := [](func(map[coordinate]*elf, coordinate) (coordinate, bool)){
		adjOccupiedFunc(coordinate{x: 0, y: -1}),
		adjOccupiedFunc(coordinate{x: 0, y: 1}),
		adjOccupiedFunc(coordinate{x: -1, y: 0}),
		adjOccupiedFunc(coordinate{x: 1, y: 0}),
	}
	turn := 0
	for {
		moves := map[coordinate][]*elf{}
		for current, elf := range elves {
			// fmt.Println(current)
			var next *coordinate
			unblocked := true
			for i := range nswe {
				direction, blocked := nswe[(turn+i)%4](elves, current)
				if blocked {
					unblocked = false
					// fmt.Println("unblocked", unblocked)
					continue
				}
				if next == nil {
					// fmt.Println("first")
					next = &direction
				}
			}
			if unblocked {
				// fmt.Println("unblocked continue")
				continue
			}
			if next == nil {
				// fmt.Println("next nil")
				continue
			}
			current.x += next.x
			current.y += next.y
			moves[current] = append(moves[current], elf)
		}
		// fmt.Println(turn, moves)
		anymove := false
		for move, candidates := range moves {
			if len(candidates) > 1 {
				continue
			}
			elf := candidates[0]
			delete(elves, elf.loc)
			elf.loc = move
			elves[move] = elf
			anymove = true
		}
		if !anymove {
			break
		}
		// draw(elves)
		turn++
	}
	return fmt.Sprintf("%v", turn), nil
}
