package main

import (
	"fmt"
	"log"
	"math"

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

type coord struct {
	x, y, z int
}

func (c coord) String() string {
	return fmt.Sprintf("%d,%d,%d", c.x, c.y, c.z)
}

func printCoords(coords []coord) {
	for i := range coords {
		fmt.Println(coords[i])
	}
}

func parse(d puzzle.Raw) []coord {
	cubeCoords := []coord{}
	lines := conversion.SplitLines(d)
	for line := range lines {
		coords := conversion.SplitCommas(lines[line])
		c := coord{}
		c.x, _ = conversion.ParseInt(coords[0])
		c.y, _ = conversion.ParseInt(coords[1])
		c.z, _ = conversion.ParseInt(coords[2])
		cubeCoords = append(cubeCoords, c)
	}
	return cubeCoords
}

func one(d puzzle.Raw) (string, error) {
	cubes := parse(d)
	total := 0
	for _, this := range cubes {
		faces := 6
		for _, next := range cubes {
			score := 0
			score += int(math.Abs(float64(this.x - next.x)))
			score += int(math.Abs(float64(this.y - next.y)))
			score += int(math.Abs(float64(this.z - next.z)))
			if score == 1 {
				faces--
			}
		}
		total += faces
	}
	return fmt.Sprintf("exposed faces %d", total), nil
}

func flood(this coord, solid, space [][][]bool) {
	space[this.x][this.y][this.z] = true
	neighbors := []coord{
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, -1},
		{0, 0, 1},
	}
	for i := range neighbors {
		neighbor := coord{
			x: this.x + neighbors[i].x,
			y: this.y + neighbors[i].y,
			z: this.z + neighbors[i].z,
		}
		if neighbor.x < 0 || neighbor.y < 0 || neighbor.z < 0 || neighbor.x > len(solid)-1 || neighbor.y > len(solid[0])-1 || neighbor.z > len(solid[1])-1 {
			// out of bounds
			continue
		}
		if space[neighbor.x][neighbor.y][neighbor.z] {
			// already flooded
			continue
		}
		if solid[neighbor.x][neighbor.y][neighbor.z] {
			// clipping
			continue
		}
		// recurse through neighborh
		flood(neighbor, solid, space)
	}
}

func spaceToCoord(space [][][]bool) []coord {
	c := []coord{}
	for x := 0; x < len(space)-1; x++ {
		for y := 0; y < len(space[x])-1; y++ {
			for z := 0; z < len(space[x][y])-1; z++ {
				if space[x][y][z] {
					c = append(c, coord{x, y, z})
				}
			}
		}
	}
	return c
}

func printSlice(b [][]bool) {
	s := "\n"
	for x := range b {
		s += "\n"
		for y := range b[x] {
			if b[x][y] {
				s += "#"
			} else {
				s += " "
			}
		}
	}
	s += "\n"
	fmt.Println(s)
}

func invert(space [][][]bool) [][][]bool {
	negative := [][][]bool{}
	for x := range space {
		negative = append(negative, [][]bool{})
		for y := range space[x] {
			negative[x] = append(negative[x], []bool{})
			for z := range space[x][y] {
				negative[x][y] = append(negative[x][y], !space[x][y][z])
			}
		}
	}
	return negative
}

func trim(space [][][]bool) [][][]bool {
	inset := [][][]bool{}
	for x := 0; x < len(space)-1; x++ {
		inset = append(inset, [][]bool{})
		for y := 0; y < len(space[x+1])-1; y++ {
			inset[x] = append(inset[x], []bool{})
			for z := 0; z < len(space[x+1][y+1])-1; z++ {
				inset[x][y] = append(inset[x][y], space[x+1][y+1][z+1])
			}
		}
	}
	return inset
}

func two(d puzzle.Raw) (string, error) {
	cubes := parse(d)
	maxXYZ := coord{0, 0, 0}
	for _, this := range cubes {
		if this.x > maxXYZ.x {
			maxXYZ.x = this.x
		}
		if this.y > maxXYZ.y {
			maxXYZ.y = this.y
		}
		if this.z > maxXYZ.z {
			maxXYZ.z = this.z
		}
	}
	fmt.Println(maxXYZ)
	solid := make([][][]bool, maxXYZ.x+2)
	space := make([][][]bool, maxXYZ.x+2)
	for x := range solid {
		solid[x] = make([][]bool, maxXYZ.y+2)
		space[x] = make([][]bool, maxXYZ.y+2)
		for y := range solid[x] {
			solid[x][y] = make([]bool, maxXYZ.z+2)
			space[x][y] = make([]bool, maxXYZ.z+2)
		}
	}
	for _, cube := range cubes {
		solid[cube.x+1][cube.y+1][cube.z+1] = true
	}
	flood(coord{0, 0, 0}, solid, space)
	fmt.Println(space)
	filled := invert(space)
	fmt.Println(filled)
	trimmed := trim(filled)
	c := spaceToCoord(trimmed)
	printCoords(c)
	total := 0
	for _, this := range c {
		faces := 6
		for _, next := range c {
			score := 0
			score += int(math.Abs(float64(this.x - next.x)))
			score += int(math.Abs(float64(this.y - next.y)))
			score += int(math.Abs(float64(this.z - next.z)))
			if score == 1 {
				faces--
			}
		}
		total += faces
	}
	return fmt.Sprintf("%d", total), nil
}
