package main

import (
	"fmt"
	"log"

	common "github.com/rbtr/go-advent-of-code/2021"
)

func main() {
	puzzle, err := common.Load()
	if err != nil {
		log.Fatal(err)
	}
	if err := one(puzzle.Sample); err != nil {
		log.Fatal(err)
	}
	if err := one(puzzle.Input); err != nil {
		log.Fatal(err)
	}
	if err := two(puzzle.Sample); err != nil {
		log.Fatal(err)
	}
	if err := two(puzzle.Input); err != nil {
		log.Fatal(err)
	}
}

func flash(oct [][]int, src common.XY) {
	oct[src.X][src.Y] = -10
	for i := 1; i >= -1; i-- {
		for j := 1; j >= -1; j-- {
			p := common.XY{src.X + i, src.Y + j}
			if p.X < 0 {
				continue
			}
			if p.X > len(oct)-1 {
				continue
			}
			if p.Y < 0 {
				continue
			}
			if p.Y > len(oct[0])-1 {
				continue
			}
			oct[p.X][p.Y]++
			if oct[p.X][p.Y] == 10 {
				flash(oct, p)
			}
		}
	}
}

func step(oct [][]int) int {
	flashes := 0

	for i := range oct {
		for j := range oct[i] {
			oct[i][j]++
		}
	}

	for i := range oct {
		for j := range oct[i] {
			if oct[i][j] >= 10 {
				flash(oct, common.XY{i, j})
			}
		}
	}

	for i := range oct {
		for j := range oct[i] {
			if oct[i][j] < 0 {
				oct[i][j] = 0
				flashes++
			}
		}
	}
	return flashes
}

func one(puzzle []byte) error {
	lines, err := common.ParseMatrix(puzzle)
	if err != nil {
		return err
	}
	count := 0
	for i := 0; i < 100; i++ {
		count += step(lines)
	}
	fmt.Printf("%d\n", count)

	return nil
}

func two(puzzle []byte) error {
	lines, err := common.ParseMatrix(puzzle)
	if err != nil {
		return err
	}
	want := len(lines) * len(lines[0])
	iter := 1
	for want != step(lines) {
		iter++
	}
	fmt.Printf("%d\n", iter)
	return nil
}
