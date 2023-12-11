package matrix

import (
	"github.com/rbtr/go-advent-of-code/pkg/conversion"
	"github.com/rbtr/go-advent-of-code/pkg/puzzle"
)

func FromData[T any](d puzzle.Data, splitRows conversion.Tokenizer, splitCols conversion.Tokenizer, parser conversion.Parser[T]) ([][]T, error) {
	var err error
	rows := splitRows(d)
	out := make([][]T, len(rows))
	for y := range rows {
		cols := splitCols(rows[y])
		out[y] = make([]T, len(cols))
		for x := range cols {
			out[x][y], err = parser(cols[x])
			if err != nil {
				return nil, err
			}
		}
	}
	return out, nil
}
