package main

import "testing"

func Test_getDistance(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{1}, 0},
		{"example2", args{12}, 3},
		{"example3", args{23}, 2},
		{"example4", args{1024}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDistance(tt.args.n); got != tt.want {
				t.Errorf("getDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
