package main

import (
	"testing"
)

func Test_partOne(t *testing.T) {
	type args struct {
		a       uint64
		b       uint64
		pairsNb int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"example1", args{65, 8921, 5}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := partOne(tt.args.a, tt.args.b, tt.args.pairsNb); gotRes != tt.wantRes {
				t.Errorf("partOne() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_partTwo(t *testing.T) {
	type args struct {
		a       uint64
		b       uint64
		pairsNb int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{"example1", args{65, 8921, 5E6}, 309},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := partTwo(tt.args.a, tt.args.b, tt.args.pairsNb); gotRes != tt.wantRes {
				t.Errorf("partTwo() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
