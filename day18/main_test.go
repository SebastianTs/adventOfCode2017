package main

import "testing"

func Test_process(t *testing.T) {
	type args struct {
		ins [][]string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"example1", args{[][]string{
			{"set", "a", "1"},
			{"add", "a", "2"},
			{"mul", "a", "a"},
			{"mod", "a", "5"},
			{"snd", "a"},
			{"set", "a", "0"},
			{"rcv", "a"},
			{"jgz", "a", "-1"},
			{"set", "a", "1"},
			{"jgz", "a", "-2"},
		}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := process(tt.args.ins); got != tt.want {
				t.Errorf("process() = %v, want %v", got, tt.want)
			}
		})
	}
}
