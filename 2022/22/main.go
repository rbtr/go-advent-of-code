package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/rbtr/go-aoc/pkg/conversion"
	"github.com/rbtr/go-aoc/pkg/puzzle"
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

type vertex struct {
	location              coordinate
	wall                  bool
	right, left, up, down *vertex
	visited               int
}

// buildLattice constructs the linked lattice and returns the entrypoint (top-left)
func buildLattice(raw [][]byte) (*vertex, map[coordinate]*vertex, coordinate) {
	sparse := map[coordinate]*vertex{}
	extents := coordinate{}
	for y := range raw {
		for x := range raw[y] {
			if strings.TrimSpace(string(raw[y][x])) == "" {
				continue
			}
			loc := coordinate{x: x, y: y}
			v := &vertex{
				location: loc,
				visited:  -1,
			}
			v.wall = string(raw[y][x]) == "#"
			sparse[loc] = v
			if x > extents.x {
				extents.x = x
			}
		}
		if y > extents.y {
			extents.y = y
		}
	}
	var start *vertex
	pos := coordinate{}
	for pos.y = 0; pos.y <= extents.y; pos.y++ {
		for pos.x = 0; pos.x <= extents.x; pos.x++ {
			current, ok := sparse[pos]
			if !ok {
				// no tile at location, continue
				continue
			}
			if start == nil {
				start = current
			}
			// tile at location
			// link right
			found := false
			rightcoord := pos
			var right *vertex
			for !found {
				rightcoord.x = (rightcoord.x + 1) % (extents.x + 1)
				right, found = sparse[rightcoord]
			}
			current.right = right
			right.left = current
			// link down
			found = false
			downcoord := pos
			var down *vertex
			for !found {
				downcoord.y = (downcoord.y + 1) % (extents.y + 1)
				down, found = sparse[downcoord]
			}
			current.down = down
			down.up = current
		}
	}
	return start, sparse, extents
}

type instruction struct {
	movement, rotation int
}

func parseInstructions(raw []byte) []instruction {
	instructions := []instruction{}
	last := 0
	for i := 0; i < len(raw); i++ {
		rot := 0
		switch string(raw[i]) {
		case "L":
			rot = -1
		case "R":
			rot = 1
		default:
			continue
		}
		move, _ := conversion.ParseInt(raw[last:i])
		instructions = append(instructions, instruction{movement: move, rotation: rot})
		last = i + 1
	}
	move, _ := conversion.ParseInt(raw[last:])
	instructions = append(instructions, instruction{movement: move, rotation: 0})
	return instructions
}

func rotate(direction, instruction int) int {
	return (((direction + instruction) % 4) + 4) % 4
}

func traverse(current *vertex, instructions []instruction) (coordinate, int) {
	direction := 0
	for _, instruction := range instructions {
		for i := 0; i < instruction.movement; i++ {
			current.visited = direction
			var next *vertex
			switch direction {
			case 0:
				next = current.right
			case 1:
				next = current.down
			case 2:
				next = current.left
			case 3:
				next = current.up
			}
			if next.wall {
				break
			}
			current = next
		}
		direction = rotate(direction, instruction.rotation)
		current.visited = direction
	}

	return coordinate{current.location.x + 1, current.location.y + 1}, direction
}

func draw(sparse map[coordinate]*vertex, extents coordinate) {
	s := "\n"
	for y := 0; y <= extents.y; y++ {
		for x := 0; x <= extents.x; x++ {
			v, ok := sparse[coordinate{x: x, y: y}]
			if !ok {
				s += " "
				continue
			}
			switch v.visited {
			case 0:
				s += ">"
			case 1:
				s += "v"
			case 2:
				s += "<"
			case 3:
				s += "^"
			default:
				if v.wall {
					s += "#"
					continue
				}
				s += "."
			}
		}
		s += "\n"
	}
	fmt.Println(s)
}

func one(d puzzle.Raw) (string, error) {
	lines := conversion.SplitLines(d)
	var lattice *vertex
	var instructions []instruction
	var sparse map[coordinate]*vertex
	var extents coordinate
	for i := range lines {
		if strings.TrimSpace(string(lines[i])) == "" {
			lattice, sparse, extents = buildLattice(lines[:i])
			instructions = parseInstructions(lines[i+1])
			break
		}
	}
	fmt.Println(instructions)
	finish, direction := traverse(lattice, instructions)
	draw(sparse, extents)
	return fmt.Sprintf("%d", 4*finish.x+1000*finish.y+direction), nil
}

func two(d puzzle.Raw) (string, error) {
	return "", nil
}
