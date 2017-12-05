package main

import (
	"testing"
)

func Test_leaveList(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[]int{0, 3, 0, 1, -3}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leaveList(tt.args.list); got != tt.want {
				t.Errorf("leaveList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_leaveListStrange(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[]int{0, 3, 0, 1, -3}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leaveListStrange(tt.args.list); got != tt.want {
				t.Errorf("leaveListStrange() = %v, want %v", got, tt.want)
			}
		})
	}
}
