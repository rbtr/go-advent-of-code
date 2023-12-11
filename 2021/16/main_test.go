package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexToBin(t *testing.T) {
	tests := []struct {
		name string
		have string
		want []int
	}{
		{
			have: "0",
			want: []int{0, 0, 0, 0},
		},
		{
			have: "D2FE28",
			want: []int{1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hexToBin(tt.have)
			assert.Equal(t, tt.want, got)
		})
	}
}

// 110100101111111000101000
// 012345678901234567890123

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		have []int
		want []packet
	}{
		{
			have: []int{1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0},
			want: []packet{
				{
					ver: 6,
					typ: 4,
				},
			},
		},
		{
			have: []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: []packet{
				{
					ver: 1,
					typ: 6,
				},
				{
					ver: 6,
					typ: 4,
				},
				{
					ver: 2,
					typ: 4,
				},
			},
		},
		{
			have: []int{1, 1, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0},
			want: []packet{
				{
					ver: 7,
					typ: 3,
				},
				{
					ver: 2,
					typ: 4,
				},
				{
					ver: 4,
					typ: 4,
				},
				{
					ver: 1,
					typ: 4,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := parse(tt.have, -1)
			assert.Equal(t, tt.want, got)
		})
	}
}
