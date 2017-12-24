package main

import (
	"reflect"
	"testing"
)

func Test_strengthOfBridge(t *testing.T) {
	type args struct {
		fn func(p, s, l int)
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"StrongestBridgeExample", args{buildBridge}, 31},
		{"LongestBridgeExample", args{buildLongestBridge}, 19},
	}
	cs = []component{
		{port0: 0, port1: 2, inUse: false},
		{port0: 2, port1: 2, inUse: false},
		{port0: 2, port1: 3, inUse: false},
		{port0: 3, port1: 4, inUse: false},
		{port0: 3, port1: 5, inUse: false},
		{port0: 0, port1: 1, inUse: false},
		{port0: 10, port1: 1, inUse: false},
		{port0: 9, port1: 10, inUse: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strengthOfBridge(tt.args.fn); got != tt.want {
				t.Errorf("strengthOfBridge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseInput(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want []component
	}{
		{"example1", args{"./input_example"}, []component{
			{port0: 0, port1: 2, inUse: false},
			{port0: 2, port1: 2, inUse: false},
			{port0: 2, port1: 3, inUse: false},
			{port0: 3, port1: 4, inUse: false},
			{port0: 3, port1: 5, inUse: false},
			{port0: 0, port1: 1, inUse: false},
			{port0: 10, port1: 1, inUse: false},
			{port0: 9, port1: 10, inUse: false},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
