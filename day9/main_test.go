package main

import "testing"

func Test_countGarbage(t *testing.T) {
	type args struct {
		stream string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"example1", args{"{}"}, 1, 0},
		{"example2", args{"{{{}}}"}, 6, 0},
		{"example3", args{"{{},{}}"}, 5, 0},
		{"example4", args{"{{{},{},{{}}}}"}, 16, 0},
		{"example5", args{"{<a>,<a>,<a>,<a>}"}, 1, 4},
		{"example6", args{"{{<ab>},{<ab>},{<ab>},{<ab>}}"}, 9, 8},
		{"example7", args{"{{<!!>},{<!!>},{<!!>},{<!!>}}"}, 9, 0},
		{"example8", args{"{{<a!>},{<a!>},{<a!>},{<ab>}}"}, 3, 17},
		{"example9", args{"<>"}, 0, 0},
		{"exampleA", args{"<random characters>"}, 0, 17},
		{"exampleB", args{"<<<<>"}, 0, 3},
		{"exampleC", args{"<{!>}>"}, 0, 2},
		{"exampleD", args{"<!!>"}, 0, 0},
		{"exampleE", args{"<!!!>>"}, 0, 0},
		{"exampleF", args{"<{o\"i!a,<{i<a>"}, 0, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := countGarbage(tt.args.stream)
			if got != tt.want {
				t.Errorf("countGarbage() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("countGarbage() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
