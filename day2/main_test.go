package main

import (
	"reflect"
	"testing"
)

func Test_rowChecksum(t *testing.T) {
	type args struct {
		row []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[]int{5, 1, 9, 5}}, 8},
		{"example2", args{[]int{7, 5, 3}}, 4},
		{"example3", args{[]int{2, 4, 6, 8}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rowChecksum(tt.args.row); got != tt.want {
				t.Errorf("rowChecksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rowChecksumEven(t *testing.T) {
	type args struct {
		row []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{[]int{5, 9, 2, 8}}, 4},
		{"example2", args{[]int{9, 4, 7, 3}}, 3},
		{"example3", args{[]int{3, 8, 6, 5}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rowChecksumEven(tt.args.row); got != tt.want {
				t.Errorf("rowChecksumEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checksum(t *testing.T) {
	type args struct {
		fn    func([]int) int
		table [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
	}{
		{"example1", args{rowChecksum, [][]int{[]int{5, 1, 9, 5}, []int{7, 5, 3}, []int{2, 4, 6, 8}}}, 18},
		{"example2", args{rowChecksumEven, [][]int{[]int{5, 9, 2, 8}, []int{9, 4, 7, 3}, []int{3, 8, 6, 5}}}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := checksum(tt.args.fn, tt.args.table); gotSum != tt.wantSum {
				t.Errorf("checksum() = %v, want %v", gotSum, tt.wantSum)
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
		want [][]int
	}{
		{"smallinput", args{"./smallinput"}, [][]int{[]int{5, 1, 9, 5}, []int{7, 5, 3}, []int{2, 4, 6, 8}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseInput(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
