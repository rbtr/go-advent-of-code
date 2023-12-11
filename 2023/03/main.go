package main

import (
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

func one(p common.Data) (string, error) {
	return "", nil
}

func two(p common.Data) (string, error) {
	return "", nil
}
