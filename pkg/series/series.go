package series

import (
	"github.com/rbtr/go-aoc/pkg/conversion"
	"github.com/rbtr/go-aoc/pkg/puzzle"
	"golang.org/x/exp/constraints"
)

func New[T any]() *builder[T] {
	return &builder[T]{}
}

type builder[T any] struct {
	d     puzzle.Raw
	split conversion.Tokenizer
	parse conversion.Parser[T]
}

func (sb *builder[T]) From(d puzzle.Raw) *builder[T] {
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

type Comparable[T comparable] interface {
	Comparable() T
}

type Set[S comparable, T Comparable[S]] map[S]struct{}

func (s Set[S, T]) Add(t T) {
	s[t.Comparable()] = struct{}{}
}

func (s Set[S, T]) Remove(t T) {
	delete(s, t.Comparable())
}

func (s Set[S, T]) Contains(t T) bool {
	_, ok := s[t.Comparable()]
	return ok
}

func (s Set[S, T]) Copy() Set[S, T] {
	copy := Set[S, T]{}
	for k := range s {
		copy[k] = struct{}{}
	}
	return copy
}
