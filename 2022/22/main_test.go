package main

import "testing"

func TestDirection(t *testing.T) {
	tests := []struct {
		name                         string
		direction, instruction, want int
	}{
		{
			name:        "up",
			direction:   0,
			instruction: 1,
			want:        1,
		},
		{
			name:        "wrap up",
			direction:   3,
			instruction: 1,
			want:        0,
		},
		{
			name:        "wrap down",
			direction:   0,
			instruction: -1,
			want:        3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate(tt.direction, tt.instruction); got != tt.want {
				t.Errorf("direction() = %v, want %v", got, tt.want)
			}
		})
	}
}
