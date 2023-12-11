package main

import (
	"fmt"

	common "github.com/rbtr/go-advent-of-code/2021"
	"go.uber.org/zap"
)

func main() {
	l, _ := zap.NewDevelopment()
	z := l.Sugar()
	puzzle, err := common.Load()
	if err != nil {
		z.Fatalw("", err)
	}
	if err := one(puzzle); err != nil {
		z.Fatalw("", err)
	}
	if err := two(puzzle); err != nil {
		z.Fatalw("", err)
	}
}

func one(puzzle *common.Puzzle) error {
	b, err := newBingo(puzzle.Input)
	if err != nil {
		return err
	}
	var winner *board
	var lastDraw int
	for _, draw := range b.draws {
		if winner != nil {
			break
		}
		if v, ok := b.field[draw]; ok {
			for _, mapper := range v {
				if winner != nil {
					break
				}
				mapper.b.rowTouches[mapper.row]++
				mapper.b.colTouches[mapper.col]++
				mapper.b.sum -= draw
				if mapper.b.rowTouches[mapper.row] == 5 || mapper.b.colTouches[mapper.col] == 5 {
					winner = mapper.b
					lastDraw = draw
				}
			}
		}
	}

	fmt.Printf("%+v,%d,%d\n", winner, lastDraw, winner.sum*lastDraw)
	return nil
}

func two(puzzle *common.Puzzle) error {
	b, err := newBingo(puzzle.Input)
	if err != nil {
		return err
	}
	var lastWinner *board
	var lastDraw int
	for _, draw := range b.draws {
		if b.boards == 0 {
			break
		}
		if v, ok := b.field[draw]; ok {
			for _, mapper := range v {
				if mapper.b.hasWon {
					continue
				}
				mapper.b.rowTouches[mapper.row]++
				mapper.b.colTouches[mapper.col]++
				mapper.b.sum -= draw
				if mapper.b.rowTouches[mapper.row] == 5 || mapper.b.colTouches[mapper.col] == 5 {
					mapper.b.hasWon = true
					lastDraw = draw
					lastWinner = mapper.b
					b.boards--
				}
			}
		}
	}

	fmt.Printf("%+v,%d,%d\n", lastWinner, lastDraw, lastWinner.sum*lastDraw)
	return nil
}
