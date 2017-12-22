package main

import (
	"testing"
)

func Test_sporifica(t *testing.T) {
	type args struct {
		node   map[[2]int]bool
		center [2]int
		bursts int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{"example1", args{map[[2]int]bool{
			[2]int{2, 0}: true,
			[2]int{0, 1}: true,
		}, [2]int{1, 1}, 10000}, 5587},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := sporifica(tt.args.node, tt.args.center, tt.args.bursts); gotCount != tt.wantCount {
				t.Errorf("sporifica() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func Test_sporificaEnhanced(t *testing.T) {
	type args struct {
		node   map[[2]int]bool
		center [2]int
		bursts int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{"example1", args{map[[2]int]bool{
			[2]int{2, 0}: true,
			[2]int{0, 1}: true,
		}, [2]int{1, 1}, 10000000}, 2511944},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := sporificaEnhanced(tt.args.node, tt.args.center, tt.args.bursts); gotCount != tt.wantCount {
				t.Errorf("sporificaEnhanced() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
