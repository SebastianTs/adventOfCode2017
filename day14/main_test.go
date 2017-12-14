package main

import "testing"

func Test_defragWithKnot(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{"example1", args{"flqrgnkx"}, 8108},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := defragWithKnot(tt.args.s); gotCount != tt.wantCount {
				t.Errorf("defragWithKnot() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
