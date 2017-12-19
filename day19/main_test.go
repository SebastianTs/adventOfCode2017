package main

import "testing"

func Test_maze(t *testing.T) {
	tests := []struct {
		name        string
		file        string
		wantVisited string
		wantSteps   int
	}{
		{"example1", "./input_example", "ABCDEF", 38},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := parseInput(tt.file)
			gotVisited, gotSteps := maze(input)
			if gotVisited != tt.wantVisited {
				t.Errorf("maze() gotVisited = %v, want %v", gotVisited, tt.wantVisited)
			}
			if gotSteps != tt.wantSteps {
				t.Errorf("maze() gotSteps = %v, want %v", gotSteps, tt.wantSteps)
			}
		})
	}
}
