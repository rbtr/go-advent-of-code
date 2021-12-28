package main

import (
	"testing"

	"github.com/rbtr/go-aoc/pkg/puzzle"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	main()
}

var p1 = puzzle.Puzzle{
	Data: []byte(`199
200
208
210
200
207
240
269
260
263
`),
	Solution: "7",
}

func TestOne(t *testing.T) {
	got, err := one(p1.Data)
	require.NoError(t, err)
	require.EqualValues(t, p1.Solution, got)
}

var p2 = puzzle.Puzzle{
	Data: []byte(`199
200
208
210
200
207
240
269
260
263`),
	Solution: "5",
}

func TestTwo(t *testing.T) {
	got, err := two(p2.Data)
	require.NoError(t, err)
	require.EqualValues(t, p2.Solution, got)
}
