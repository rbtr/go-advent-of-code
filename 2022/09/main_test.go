package main

import (
	"reflect"
	"testing"
)

func TestMove(t *testing.T) {
	tests := []struct {
		name string
		a, b coord
		want coord
	}{
		{
			name: "linear",
			a:    coord{x: 2, y: 0},
			b:    coord{x: 0, y: 0},
			want: coord{x: 1, y: 0},
		},
		{
			name: "neg linear",
			a:    coord{x: -2, y: 0},
			b:    coord{x: 0, y: 0},
			want: coord{x: -1, y: 0},
		},
		{
			name: "diag",
			a:    coord{x: -2, y: 1},
			b:    coord{x: 0, y: 0},
			want: coord{x: -1, y: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMove(tt.a, tt.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("move() = %v, want %v", got, tt.want)
			}
		})
	}
}
