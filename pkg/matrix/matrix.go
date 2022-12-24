package matrix

import (
	"github.com/rbtr/go-aoc/pkg/conversion"
	"github.com/rbtr/go-aoc/pkg/puzzle"
)

type IndexedParser[T any] func(int, int, []byte) (T, error)

func FromData[T any](d puzzle.Raw, splitRows conversion.Tokenizer, splitCols conversion.Tokenizer, parser IndexedParser[T]) ([][]T, error) {
	var err error
	rows := splitRows(d)
	out := make([][]T, len(rows))
	for y := range rows {
		cols := splitCols(rows[y])
		out[y] = make([]T, len(cols))
		for x := range cols {
			out[y][x], err = parser(y, x, cols[x])
			if err != nil {
				return nil, err
			}
		}
	}
	return out, nil
}
