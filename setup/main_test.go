package main

import (
	"testing"

	"github.com/rbtr/go-advent-of-code/pkg/puzzle"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	main()
}

var p1 = puzzle.Puzzle{
	Data: []byte(`
`),
	Solution: "",
}

func TestOne(t *testing.T) {
	got, err := one(p1.Data)
	require.NoError(t, err)
	require.EqualValues(t, p1.Solution, got)
}

var p2 = puzzle.Puzzle{
	Data: []byte(`
`),
	Solution: "",
}

func TestTwo(t *testing.T) {
	got, err := two(p2.Data)
	require.NoError(t, err)
	require.EqualValues(t, p2.Solution, got)
}
