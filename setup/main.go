package main

import (
	"log"

	"github.com/rbtr/go-advent-of-code/pkg/parse"
	"github.com/rbtr/go-advent-of-code/pkg/puzzle"
)

func main() {
	in, err := parse.Input()
	if err != nil {
		log.Fatal(err)
	}
	if err := puzzle.Run(in, one, two); err != nil {
		log.Fatal(err)
	}
}

func one(d puzzle.Data) (string, error) {
	return "", nil
}

func two(d puzzle.Data) (string, error) {
	return "", nil
}
