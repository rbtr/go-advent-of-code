package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoll(t *testing.T) {
	tests := []struct {
		name string
		turn int
		want int
	}{
		{
			turn: 0,
			want: 1 + 2 + 3,
		},
		{
			turn: 1,
			want: 4 + 5 + 6,
		},
		{
			turn: 2,
			want: 7 + 8 + 9,
		},
		{
			turn: 3,
			want: 10 + 1 + 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := roll(10, tt.turn)
			assert.Equal(t, tt.want, out)
		})
	}
}

func TestMove(t *testing.T) {
	tests := []struct {
		name    string
		current int
		move    int
		want    int
	}{
		{
			current: 8,
			move:    3,
			want:    1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := move(10, tt.current, tt.move)
			assert.Equal(t, tt.want, out)
		})
	}
}
