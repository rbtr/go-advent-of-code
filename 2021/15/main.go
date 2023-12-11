package main

import (
	"container/heap"
	"fmt"
	"log"

	common "github.com/rbtr/go-advent-of-code/2021"
)

func main() {
	puzzle, err := common.Load()
	if err != nil {
		log.Fatal(err)
	}
	common.Solve(puzzle, one, two)
}

func reduce(matr [][]int) int {
	for i := 0; i < len(matr); i++ {
		for j := 0; j < len(matr[i]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			up := -1
			if i > 0 {
				up = matr[i-1][j]
			}
			left := -1
			if j > 0 {
				left = matr[i][j-1]
			}
			if up < 0 {
				up = left
			}
			if left < 0 {
				left = up
			}
			if up < left {
				matr[i][j] = matr[i][j] + up
			} else {
				matr[i][j] = matr[i][j] + left
			}
		}
	}
	return matr[len(matr)-1][len(matr[len(matr)-1])-1]
}

func one(puzzle common.Data) (string, error) {
	matr, err := common.ParseMatrix(puzzle)
	matr[0][0] = 0
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", reduce(matr)), nil
}

func tile(in [][]int, times int) [][]int {
	out := make([][]int, len(in)*times)
	for i := range out {
		ii := i % len(in)
		out[i] = make([]int, len(in[ii])*times)
		for j := range out[i] {
			jj := j % len(in)
			out[i][j] = in[ii][jj] + (i / len(in)) + (j / len(in[ii]))
			if (i/len(in)) > 0 || (j/len(in[ii])) > 0 {
				if out[i][j] > 9 {
					out[i][j] = out[i][j] % 9
				}
			}
		}
	}
	return out
}

type XY common.XY

func build(matr [][]int) ([][]*Vert, *priorityq) {
	max := len(matr) * len(matr[0]) * 10
	out := make([][]*Vert, len(matr))
	unvisited := make(priorityq, 0)
	heap.Init(&unvisited)
	for i := range matr {
		out[i] = make([]*Vert, len(matr[i]))
		for j := range matr[i] {
			v := &Vert{coord: XY{j, i}, totalcost: max, entrycost: matr[i][j]}
			out[i][j] = v
			heap.Push(&unvisited, v)
		}
	}
	return out, &unvisited
}

func visit(current *Vert, grid [][]*Vert, unvisited *priorityq) {
	nexts := []XY{
		{ // right
			X: current.coord.X + 1,
			Y: current.coord.Y,
		},
		{ // left
			X: current.coord.X - 1,
			Y: current.coord.Y,
		},
		{ // down
			X: current.coord.X,
			Y: current.coord.Y + 1,
		},
		{ // up
			X: current.coord.X,
			Y: current.coord.Y - 1,
		},
	}

	for _, next := range nexts {
		if next.Y < 0 || next.X < 0 || next.Y > len(grid)-1 || next.X > len(grid[next.Y])-1 {
			continue
		}

		nextvert := grid[next.Y][next.X]
		pathcost := current.totalcost + nextvert.entrycost
		if nextvert.totalcost < 0 || pathcost < nextvert.totalcost {
			nextvert.totalcost = pathcost
		}
		unvisited.update(nextvert)
	}
}

func two(puzzle common.Data) (string, error) {
	matr, err := common.ParseMatrix(puzzle)
	matr = tile(matr, 5)
	matr[0][0] = 0
	if err != nil {
		return "", err
	}

	grid, pq := build(matr)
	grid[0][0].totalcost = 0
	pq.update(grid[0][0])
	for pq.Len() > 0 {
		vert := heap.Pop(pq).(*Vert)
		visit(vert, grid, pq)
	}
	return fmt.Sprintf("total cost %d", grid[len(grid)-1][len(grid[len(grid)-1])-1].totalcost), nil
}
