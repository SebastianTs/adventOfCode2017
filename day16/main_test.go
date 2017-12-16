package main

import (
	"reflect"
	"testing"
)

func Test_parseInput(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"ParseAllCases", args{"./input_example"}, []string{"s1", "x0/1", "pa/b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dance(t *testing.T) {
	var progs = [16]rune{}
	for i := 0; i < len(progs); i++ {
		progs[i] = rune('a' + i)
	}
	type args struct {
		progs *[16]rune
		moves []string
	}
	tests := []struct {
		name string
		args args
	}{
		{"ExampleDance", args{&progs, []string{"s1", "x0/1", "pa/b"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dance(tt.args.progs, tt.args.moves)
			if progs[0] != 'b' {
				t.Fail()
			}
		})
	}
}

func Test_dancers(t *testing.T) {
	var progs = [16]rune{}
	for i := 0; i < len(progs); i++ {
		progs[i] = rune('a' + i)
	}
	type args struct {
		progs *[16]rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"AtoF", args{&progs}, "abcdefghijklmnop"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dancers(tt.args.progs); got != tt.want {
				t.Errorf("dancers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_billionDances(t *testing.T) {
	var progs = [16]rune{}
	for i := 0; i < len(progs); i++ {
		progs[i] = rune('a' + i)
	}
	type args struct {
		progs *[16]rune
		moves []string
	}
	tests := []struct {
		name string
		args args
	}{
		{"BillionExampleDances", args{&progs, []string{"s1", "x0/1", "pa/b"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if progs[0] != 'a' {
				t.Fail()
			}
		})
	}
}
