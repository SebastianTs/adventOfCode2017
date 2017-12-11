package main

import "testing"

func Test_distance(t *testing.T) {
	type args struct {
		steps []string
	}
	tests := []struct {
		name     string
		args     args
		wantDist int
		wantMax  int
	}{
		{"example1", args{[]string{"ne", "ne", "ne"}}, 3, 3},
		{"example2", args{[]string{"ne", "ne", "sw", "sw"}}, 0, 2},
		{"example3", args{[]string{"ne", "ne", "s", "s"}}, 2, 2},
		{"example4", args{[]string{"se", "sw", "se", "sw", "sw"}}, 3, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDist, gotMax := distance(tt.args.steps)
			if gotDist != tt.wantDist {
				t.Errorf("distance() gotDist = %v, want %v", gotDist, tt.wantDist)
			}
			if gotMax != tt.wantMax {
				t.Errorf("distance() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
