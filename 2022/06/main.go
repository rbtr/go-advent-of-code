package main

import (
	"log"
	"strconv"

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
	start := 3
	for ; start < len(d); start++ {
		m := map[byte]struct{}{
			d[start-3]: {},
			d[start-2]: {},
			d[start-1]: {},
			d[start]:   {},
		}
		if len(m) == 4 {
			break
		}
	}
	return strconv.Itoa(start + 1), nil
}

func two(d puzzle.Raw) (string, error) {
	start := 14
	for ; start < len(d); start++ {
		m := map[byte]struct{}{}
		for j := 0; j < 14; j++ {
			m[d[start-j]] = struct{}{}
		}
		if len(m) == 14 {
			break
		}
	}
	return strconv.Itoa(start + 1), nil
}
