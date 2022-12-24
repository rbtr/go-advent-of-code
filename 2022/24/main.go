package main

import (
	"fmt"
	"log"

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

type blizzard struct {
	direction coordinate
}

func parse(d puzzle.Raw) (coordinate, map[coordinate][]*blizzard) {
	sparse := map[coordinate][]*blizzard{}
	lines := conversion.SplitLines(d)
	for i := 0; i < len(lines); i++ {
		chars := conversion.SplitCharacters(lines[i])
		for j := 0; j < len(chars); j++ {
			char := string(chars[j])
			if char == "." {
				continue
			}
			b := blizzard{}
			switch char {
			case "<":
				b.direction.x = -1
			case ">":
				b.direction.x = 1
			case "^":
				b.direction.y = -1
			case "v":
				b.direction.y = 1
			case "#":
			default:
				fmt.Println("unrecognized blizzard", char)
			}
			coord := coordinate{x: j, y: i}
			sparse[coord] = append(sparse[coord], &b)
		}
	}
	return coordinate{x: len(lines[0]), y: len(lines)}, sparse
}

func advance(extents coordinate, sparse map[coordinate][]*blizzard) map[coordinate][]*blizzard {
	next := map[coordinate][]*blizzard{}
	for coord, blizzards := range sparse {
		for _, blizzard := range blizzards {
			nextPos := coord
			nextPos.x += blizzard.direction.x
			nextPos.y += blizzard.direction.y
			if blizzard.direction.x != 0 || blizzard.direction.y != 0 {
				nextPos.x = (((nextPos.x-1)%(extents.x-2))+(extents.x-2))%(extents.x-2) + 1
				nextPos.y = (((nextPos.y-1)%(extents.y-2))+(extents.y-2))%(extents.y-2) + 1
			}
			next[nextPos] = append(next[nextPos], blizzard)
		}
	}
	return next
}

func draw(extents coordinate, sparse map[coordinate][]*blizzard) {
	s := "\n"
	for i := 0; i < extents.y; i++ {
		for j := 0; j < extents.x; j++ {
			blizzards, ok := sparse[coordinate{x: j, y: i}]
			if !ok {
				s += "."
				continue
			}
			if len(blizzards) > 1 {
				s += fmt.Sprintf("%d", len(blizzards))
				continue
			}
			switch blizzards[0].direction {
			case coordinate{x: -1, y: 0}:
				s += "<"
			case coordinate{x: 1, y: 0}:
				s += ">"
			case coordinate{x: 0, y: -1}:
				s += "^"
			case coordinate{x: 0, y: 1}:
				s += "v"
			case coordinate{}:
				s += "#"
			}
		}
		s += "\n"
	}
	fmt.Println(s)
}

func precompute(extents coordinate, turns int, initial map[coordinate][]*blizzard) []map[coordinate][]*blizzard {
	frames := []map[coordinate][]*blizzard{initial}
	for i := 0; i < turns; i++ {
		frames = append(frames, advance(extents, frames[i]))
	}
	return frames
}

// func traverse(position, end, extents coordinate, best, turn, max int, frames []map[coordinate][]*blizzard) int {
// 	// fmt.Println(turn, best)
// 	if turn >= max {
// 		return math.MaxInt
// 	}
// 	// if our position is occupied by a blizzard in this turn, die
// 	if _, ok := frames[turn][position]; ok {
// 		return math.MaxInt
// 	}
// 	// if our position is the end position, done
// 	if position == end {
// 		// fmt.Println("end", best)
// 		return turn
// 	}
// 	options := []coordinate{}
// 	// if our position is the start position, we can only wait or move south
// 	if position != (coordinate{x: 0, y: -1}) {
// 		// append west, east, north
// 		options = append(options, []coordinate{{x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: -1}}...)
// 	}
// 	// append south, wait
// 	options = append(options, []coordinate{{x: 0, y: 1}, {x: 0, y: 0}}...)

// 	for _, next := range options {
// 		// recurse over each available option
// 		next.x += position.x
// 		next.y += position.y
// 		if next != (coordinate{x: 0, y: -1}) {
// 			if next.x < 0 || next.x > extents.x {
// 				continue
// 			}
// 			if next.y < 0 || next.y > extents.y {
// 				continue
// 			}
// 		}
// 		result := traverse(next, end, extents, best, turn+1, max, frames)
// 		if result < best {
// 			fmt.Println(best)
// 			best = result
// 		}
// 	}
// 	// otherwise, wait or move in any any direction
// 	return best
// }

type node struct {
	position     coordinate
	frame, score int
	neighbors    []*node
	valid        bool
}

func cloud(extents coordinate, frames []map[coordinate][]*blizzard) [][][]*node {
	nodes := [][][]*node{}
	for z := 0; z < len(frames)-1; z++ {
		nodes = append(nodes, [][]*node{})
		for y := 0; y < extents.y+2; y++ {
			nodes[z] = append(nodes[z], []*node{})
			for x := 0; x < extents.x+2; x++ {
				pos := coordinate{x: x, y: y}
				_, occupied := frames[z][pos]
				this := &node{valid: !occupied, position: coordinate{x: x, y: y}, frame: z, score: -1, neighbors: []*node{}}
				nodes[z][y] = append(nodes[z][y], this)
			}
		}
	}
	return nodes
}

func link(extents coordinate, cloud [][][]*node) *node {
	options := []coordinate{{x: -1, y: 0}, {x: 1, y: 0}, {x: 0, y: -1}, {x: 0, y: 1}, {x: 0, y: 0}}
	for z := 0; z < len(cloud)-1; z++ {
		for y := 0; y < len(cloud[z]); y++ {
			for x := 0; x < len(cloud[z][y]); x++ {
				// fmt.Println(z, y, x, extents)
				this := cloud[z][y][x]
				if !this.valid {
					continue
				}
				for _, nextpos := range options {
					nextpos.x += x
					nextpos.y += y
					if nextpos.y < 0 || nextpos.x < 0 {
						continue
					}
					if nextpos.y > extents.y || nextpos.x > extents.x {
						continue
					}
					next := cloud[z+1][nextpos.y][nextpos.x]
					if !next.valid {
						continue
					}
					this.neighbors = append(this.neighbors, next)
				}
			}
		}
	}
	return cloud[0][0][1]
}

func search(current *node, finish coordinate) {
	// fmt.Println(current.position, current.neighbors)
	// fmt.Println(current.neighbors)
	if current.position.x == finish.x && current.position.y == finish.y {
		return
	}
	for _, next := range current.neighbors {
		if next.score >= 0 && next.score <= current.score+1 {
			continue
		}
		next.score = current.score + 1
		search(next, finish)
	}
}

func one(d puzzle.Raw) (string, error) {
	extents, blizzards := parse(d)
	// fmt.Println(extents)
	// turn := 0
	// for {
	// 	if turn == 3 {
	// 		break
	// 	}
	// 	draw(extents, blizzards)
	// 	blizzards = advance(extents, blizzards)
	// 	turn++
	// }
	max := 1000 // int(math.Pow(float64(extents.x+extents.y), 3))
	frames := precompute(extents, max, blizzards)
	cloud := cloud(extents, frames)
	root := link(extents, cloud)
	root.score = 0
	finish := coordinate{x: extents.x - 2, y: extents.y - 1}
	search(root, finish)
	turn := 0
	for turn = range cloud {
		if cloud[turn][finish.y][finish.x].score > 0 {
			break
		}
	}
	return fmt.Sprintf("best: %d", turn), nil
}

func two(d puzzle.Raw) (string, error) {
	return "", nil
}
