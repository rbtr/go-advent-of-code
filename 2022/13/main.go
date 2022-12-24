package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rbtr/go-aoc/pkg/conversion"
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

func compare(left, right []any) bool {
	for i := range left {
		for j := range right {
			if l, ok := (left[i]).([]any); ok {
				if r, ok := (right[j]).([]any); ok {
					if !compare(l, r) {
						return false
					}
				}
			}
		}
	}
	return false
}

func one(d puzzle.Raw) (string, error) {
	lines := conversion.SplitLines(d)
	for i := 0; i < len(lines)-1; i += 2 {
		var left, right []any
		_ = json.Unmarshal(lines[i], &left)
		_ = json.Unmarshal(lines[i+1], &right)
		fmt.Println(left, right)
	}
	return "", nil
}

func two(d puzzle.Raw) (string, error) {
	return "", nil
}
