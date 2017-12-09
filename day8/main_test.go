package main

import "testing"

func Test_processInstructionSet(t *testing.T) {
	type args struct {
		in [][]string
	}
	tests := []struct {
		name      string
		args      args
		wantMax   int
		wantFinal int
	}{
		{"example1", args{[][]string{{"b", "inc", "5", "if", "a", ">", "1"}, {"a", "inc", "1", "if", "b", "<", "5"}, {"c", "dec", "-10", "if", "a", ">=", "1"}, {"c", "inc", "-20", "if", "c", "==", "10"}}}, 1, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax, gotFinal := processInstructionSet(tt.args.in)
			if gotMax != tt.wantMax {
				t.Errorf("processInstructionSet() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
			if gotFinal != tt.wantFinal {
				t.Errorf("processInstructionSet() gotFinal = %v, want %v", gotFinal, tt.wantFinal)
			}
		})
	}
}
