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

func one(d puzzle.Raw) (string, error) {
	s, err := series.New[int]().From(d).Split(conversion.SplitLines).Parse(conversion.ParseInt).Build()
	if err != nil {
		return "", err
	}
	sum := 0
	for i := range s {
		sum += (s[i] / 3) - 2
	}
	return strconv.Itoa(sum), nil
}

func two(d puzzle.Raw) (string, error) {
	s, err := series.New[int]().From(d).Split(conversion.SplitLines).Parse(conversion.ParseInt).Build()
	if err != nil {
		return "", err
	}
	for i := range s {
		fuels := []int{}
		mass := s[i]
		for {
			delta := (mass / 3) - 2
			if delta <= 0 {
				break
			}
			mass = delta
			fuels = append(fuels, delta)
		}
		s[i] = series.Sum(fuels)
	}
	return strconv.Itoa(series.Sum(s)), nil
}
