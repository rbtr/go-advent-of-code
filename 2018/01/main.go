package main

import (
	"log"
	"strconv"

	"github.com/rbtr/go-aoc/pkg/conversion"
	"github.com/rbtr/go-aoc/pkg/parse"
	"github.com/rbtr/go-aoc/pkg/puzzle"
	"github.com/rbtr/go-aoc/pkg/series"
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
	s, err := series.New[int]().From(d).Split(conversion.SplitLines).Parse(conversion.ParseSignedInts).Build()
	if err != nil {
		return "", err
	}
	freq := 0
	for i := range s {
		freq += s[i]
	}
	return strconv.Itoa(freq), nil
}

func two(d puzzle.Data) (string, error) {
	s, err := series.New[int]().From(d).Split(conversion.SplitLines).Parse(conversion.ParseSignedInts).Build()
	if err != nil {
		return "", err
	}
	freq := 0
	freqs := map[int]any{}
	for i := 0; ; i++ {
		freq += s[i%len(s)]
		if _, seen := freqs[freq]; seen {
			break
		}
		freqs[freq] = nil
	}
	return strconv.Itoa(freq), nil
}
