package main

import (
	"testing"
)

var exampleMap = map[int]int{
	0: 3,
	1: 2,
	4: 4,
	6: 4,
}

func Test_firewallStateDelay(t *testing.T) {
	type args struct {
		m map[int]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{exampleMap}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firewallStateDelay(tt.args.m); got != tt.want {
				t.Errorf("firewallStateDelay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firewallState(t *testing.T) {
	type args struct {
		m map[int]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{exampleMap}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firewallState(tt.args.m); got != tt.want {
				t.Errorf("firewallState() = %v, want %v", got, tt.want)
			}
		})
	}
}
