package series

import (
	"constraints"

	"github.com/rbtr/go-aoc/pkg/conversion"
	"github.com/rbtr/go-aoc/pkg/puzzle"
)

func New[T any]() *builder[T] {
	return &builder[T]{}
}

type builder[T any] struct {
	d     puzzle.Data
	split conversion.Tokenizer
	parse conversion.Parser[T]
}

func (sb *builder[T]) From(d puzzle.Data) *builder[T] {
	sb.d = d
	return sb
}

func (sb *builder[T]) Split(delim conversion.Tokenizer) *builder[T] {
	sb.split = delim
	return sb
}

func (sb *builder[T]) Parse(parser conversion.Parser[T]) *builder[T] {
	sb.parse = parser
	return sb
}

func (sb *builder[T]) Build() ([]T, error) {
	var err error
	toks := sb.split(sb.d)
	out := make([]T, len(toks))
	for i := range toks {
		out[i], err = sb.parse(toks[i])
		if err != nil {
			return nil, err
		}
	}
	return out, nil
}

func Sum[T constraints.Ordered](in []T) T {
	var out T
	for i := range in {
		out += in[i]
	}
	return out
}
