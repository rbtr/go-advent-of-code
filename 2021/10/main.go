package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"sort"

	common "github.com/rbtr/go-advent-of-code/2021"
)

func main() {
	puzzle, err := common.Load()
	if err != nil {
		log.Fatal(err)
	}
	if err := one(puzzle); err != nil {
		log.Fatal(err)
	}
	if err := two(puzzle); err != nil {
		log.Fatal(err)
	}
}

var closers = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var errVals = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func one(puzzle *common.Puzzle) error {
	sc := bufio.NewScanner(bytes.NewReader(puzzle.Input))
	sc.Split(bufio.ScanLines)

	bad := 0
	for sc.Scan() {
		closes := []rune{}
		line := sc.Text()
		for _, r := range line {
			closer, ok := closers[r]
			if ok {
				// r has closer so it is an opener, push closer
				closes = append(closes, closer)
				continue
			}
			// r has no closer, so it is a closer
			// is it the one we need?
			if r == closes[len(closes)-1] {
				// yes
				// pop it
				closes = closes[:len(closes)-1]
				continue
			}

			// no
			// break
			bad += errVals[r]
			break
		}
	}
	fmt.Printf("1: %d\n", bad)
	return nil
}

var compVals = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func two(puzzle *common.Puzzle) error {
	sc := bufio.NewScanner(bytes.NewReader(puzzle.Input))
	sc.Split(bufio.ScanLines)

	scores := []int{}
	for sc.Scan() {
		closes := []rune{}
		line := sc.Text()
		bad := false
		for _, r := range line {
			closer, ok := closers[r]
			if ok {
				// r has closer so it is an opener, push closer
				closes = append(closes, closer)
				continue
			}
			// r has no closer, so it is a closer
			// is it the one we need?
			if r == closes[len(closes)-1] {
				// yes
				// pop it
				closes = closes[:len(closes)-1]
				continue
			}

			// no
			// break
			bad = true
			break
		}
		if bad {
			continue
		}
		score := 0
		for i := len(closes) - 1; i >= 0; i-- {
			score = score * 5
			score += compVals[closes[i]]
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	fmt.Printf("2: %d\n", scores[len(scores)/2])
	return nil
}
