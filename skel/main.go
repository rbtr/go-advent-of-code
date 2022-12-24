package main

import (
	"log"

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

func one(d puzzle.Raw) (string, error) {
	return "", nil
}

func two(d puzzle.Raw) (string, error) {
	return "", nil
}
