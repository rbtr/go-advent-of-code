package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/rbtr/go-aoc/pkg/constants"
	"github.com/rbtr/go-aoc/pkg/conversion"
	"github.com/rbtr/go-aoc/pkg/matrix"
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

type coordinate struct {
	row, col int
}

type node struct {
	coordinate coordinate
	elevation  int
	score      int
	reachable  series.Set[coordinate, *node]
	reverse    series.Set[coordinate, *node]
}

func (this *node) Comparable() coordinate {
	return this.coordinate
}

func (this *node) isReachable(next *node) bool {
	// if the next node is equal or below us, consider it reachable
	if this.elevation >= next.elevation {
		return true
	}
	// if the next node is at most one above us, consider it reachable
	return next.elevation-this.elevation <= 1
}

func printMaze(maze [][]*node) {
	// time.Sleep(250 * time.Millisecond)
	fmt.Print("\033[H\033[2J")
	s := ""
	for i := range maze {
		for j := range maze[i] {
			s += fmt.Sprintf("%4d", maze[i][j].score)
		}
		s += "\n"
	}
	s += "\n"
	fmt.Println(s)
	// _ = os.WriteFile("out", []byte(s), 0o644)
}

// djikstras
// for every "reachable" node, jump to that node, set its reachablity score
// after an exhaustive walk of the maze, the end should hold the lowest possible route score
func search(maze [][]*node, current, end coordinate) {
	// printMaze(maze)
	if current == end {
		// fmt.Println("end")
		return
	}
	currentNode := maze[current.row][current.col]
	for next := range currentNode.reachable {
		nextNode := maze[next.row][next.col]
		if nextNode.score >= 0 && (currentNode.score+1) >= nextNode.score {
			continue
		}
		nextNode.score = currentNode.score + 1
		search(maze, next, end)
	}
}

func graph(m [][]*node) {
	for row := range m {
		for col := range m[row] {
			this := m[row][col]
			for _, delta := range []coordinate{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				if row+delta.row < 0 || row+delta.row > len(m)-1 {
					continue
				}
				if col+delta.col < 0 || col+delta.col > len(m[row])-1 {
					continue
				}
				next := m[row+delta.row][col+delta.col]
				if this.isReachable(next) {
					this.reachable.Add(next)
					next.reverse.Add(this)
				}
			}
		}
	}
}

func one(d puzzle.Raw) (string, error) {
	var start, end coordinate
	maze, _ := matrix.FromData(
		d,
		conversion.SplitLines,
		conversion.SplitCharacters,
		func(row, col int, b []byte) (*node, error) {
			elevationString := string(b)
			node := &node{
				score:     -1,
				reachable: series.Set[coordinate, *node]{},
				reverse:   series.Set[coordinate, *node]{},
			}
			switch elevationString {
			case "S":
				start.row, start.col = row, col
				elevationString = "a"
				node.score = 0
			case "E":
				end.row, end.col = row, col
				elevationString = "z"
			}
			node.coordinate = coordinate{row, col}
			node.elevation = strings.Index(constants.Lower, elevationString)
			return node, nil
		},
	)
	graph(maze)
	search(maze, start, end)
	return fmt.Sprintf("shortest %d", maze[end.row][end.col].score), nil
}

func reverseSearch(maze [][]*node, current coordinate) {
	currentNode := maze[current.row][current.col]
	if currentNode.elevation == 0 {
		return
	}
	for next := range currentNode.reverse {
		nextNode := maze[next.row][next.col]
		if nextNode.score >= 0 && (currentNode.score+1) >= nextNode.score {
			continue
		}
		nextNode.score = currentNode.score + 1
		reverseSearch(maze, next)
	}
}

func two(d puzzle.Raw) (string, error) {
	var start, end coordinate
	maze, _ := matrix.FromData(
		d,
		conversion.SplitLines,
		conversion.SplitCharacters,
		func(row, col int, b []byte) (*node, error) {
			elevationString := string(b)
			node := &node{
				score:     -1,
				reachable: series.Set[coordinate, *node]{},
				reverse:   series.Set[coordinate, *node]{},
			}
			switch elevationString {
			case "S":
				start.row, start.col = row, col
				elevationString = "a"
			case "E":
				end.row, end.col = row, col
				elevationString = "z"
				node.score = 0
			}
			node.coordinate = coordinate{row, col}
			node.elevation = strings.Index(constants.Lower, elevationString)
			return node, nil
		},
	)
	graph(maze)
	reverseSearch(maze, end)
	// printMaze(maze)
	closest := math.MaxInt
	for row := range maze {
		for col := range maze[row] {
			this := maze[row][col]
			if this.elevation == 0 && this.score >= 0 {
				if this.score < closest {
					closest = this.score
				}
			}
		}
	}
	return fmt.Sprintf("closest %d", closest), nil
}
