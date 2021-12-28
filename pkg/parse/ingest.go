package parse

import (
	"os"

	"github.com/rbtr/go-aoc/pkg/puzzle"
)

const (
	sample = "sample"
	input  = "input"
)

func ReadData(name string) (puzzle.Data, error) {
	return os.ReadFile(name)
}

func Input() (puzzle.Data, error) {
	return ReadData(input)
}
