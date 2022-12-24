package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/rbtr/go-aoc/pkg/conversion"
	"github.com/rbtr/go-aoc/pkg/matrix"
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

type tree struct {
	height  int
	view    int
	visible bool
}

func parseTree(in []byte) (*tree, error) {
	height, err := strconv.Atoi(string(in))
	if err != nil {
		return nil, err
	}
	return &tree{height: height}, nil
}

func one(d puzzle.Raw) (string, error) {
	grid, _ := matrix.FromData(d, conversion.SplitLines, conversion.SplitCharacters, parseTree)

	for row := range grid {
		tallest := -1
		for col := range grid[row] {
			this := grid[row][col]
			if this.height > tallest {
				this.visible = true
				tallest = this.height
			}
		}
		tallest = -1
		for col := len(grid[row]) - 1; col >= 0; col-- {
			this := grid[row][col]
			if this.height > tallest {
				this.visible = true
				tallest = this.height
			}
		}
	}
	for col := range grid[0] {
		tallest := -1
		for row := range grid {
			this := grid[row][col]
			if this.height > tallest {
				this.visible = true
				tallest = this.height
			}
		}
		tallest = -1
		for row := len(grid) - 1; row >= 0; row-- {
			this := grid[row][col]
			if this.height > tallest {
				this.visible = true
				tallest = this.height
			}
		}
	}

	total := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col].visible {
				total++
			}
		}
	}

	return fmt.Sprintf("%d", total), nil
}

func two(d puzzle.Raw) (string, error) {
	grid, _ := matrix.FromData(d, conversion.SplitLines, conversion.SplitCharacters, parseTree)

	for row := range grid {
		for col := range grid[row] {
			this := grid[row][col]
			// look left
			left := 0
			// tallest := -1
			for i := col - 1; i >= 0; i-- {
				sighted := grid[row][i]
				// fmt.Println(this.height, sighted.height)
				left++
				if sighted.height >= this.height {
					break
				}
				// if sighted.height >= tallest {
				// 	left++
				// 	tallest = sighted.height
				// }
			}
			// fmt.Println("left", left)
			// right
			right := 0
			// tallest = -1
			for i := col + 1; i < len(grid[row]); i++ {
				sighted := grid[row][i]
				// fmt.Println(this.height, sighted.height)
				right++
				if sighted.height >= this.height {
					break
				}
				// if sighted.height >= tallest {
				// 	right++
				// 	tallest = sighted.height
				// }
			}
			// fmt.Println("right", right)
			// up
			up := 0
			// tallest = -1
			for i := row - 1; i >= 0; i-- {
				sighted := grid[i][col]
				// fmt.Println(this.height, sighted.height)
				up++
				if sighted.height >= this.height {
					break
				}
				// if sighted.height >= tallest {
				// 	up++
				// 	tallest = sighted.height
				// }
			}
			// fmt.Println("up", up)
			// down
			down := 0
			// tallest = -1
			for i := row + 1; i < len(grid); i++ {
				sighted := grid[i][col]
				// fmt.Println(this.height, sighted.height)
				down++
				if sighted.height >= this.height {
					break
				}
				// if sighted.height >= tallest {
				// 	down++
				// 	tallest = sighted.height
				// }
			}
			// fmt.Println("down", down)
			// fmt.Println(this.height, left, right, up, down)
			this.view = left * right * up * down
		}
	}
	best := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col].view > best {
				best = grid[row][col].view
			}
		}
	}
	return fmt.Sprintf("%d", best), nil
}
