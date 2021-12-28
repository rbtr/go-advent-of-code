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
	s, err := series.FromData(d, conversion.SplitLines, conversion.ParseInt)
	if err != nil {
		return "", err
	}
	sum := 0
	for i := 1; i < len(s); i++ {
		if s[i-1] < s[i] {
			sum++
		}
	}
	return strconv.Itoa(sum), nil
}

func two(d puzzle.Data) (string, error) {
	s, err := series.FromData(d, conversion.SplitLines, conversion.ParseInt)
	if err != nil {
		return "", err
	}
	sum := 0
	for i := 1; i < len(s)-2; i++ {
		if s[i-1]+s[i]+s[i+1] < s[i]+s[i+1]+s[i+2] {
			sum++
		}
	}
	return strconv.Itoa(sum), nil
}
