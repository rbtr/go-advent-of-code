package series

import (
	"github.com/rbtr/go-aoc/pkg/conversion"
	"github.com/rbtr/go-aoc/pkg/puzzle"
)

func FromData[T any](d puzzle.Data, split conversion.Tokenizer, parser conversion.Parser[T]) ([]T, error) {
	var err error
	toks := split(d)
	out := make([]T, len(toks))
	for i := range toks {
		out[i], err = parser(toks[i])
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}
