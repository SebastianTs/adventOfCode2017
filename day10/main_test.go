package main

import (
	"reflect"
	"testing"
)

func Test_revSubList(t *testing.T) {
	type args struct {
		s    int
		l    int
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Odd", args{0, 5, []int{0, 1, 2, 3, 4}}, []int{4, 3, 2, 1, 0}},
		{"Even", args{0, 6, []int{0, 1, 2, 3, 4, 5}}, []int{5, 4, 3, 2, 1, 0}},
		{"Circle", args{4, 3, []int{0, 1, 2, 3, 4, 5}}, []int{4, 1, 2, 3, 0, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := revSubList(tt.args.s, tt.args.l, tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("revSubList() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func Test_knot(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"example1", args{[]int{3, 4, 1, 5, 0}}, []int{3, 4, 2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knot(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("knot() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
func Test_sToDenseShort(t *testing.T) {
	type args struct {
		sparse []byte
	}
	tests := []struct {
		name      string
		args      args
		wantDense byte
	}{
		{"short", args{[]byte{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22}}, byte(64)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDense := sToDenseShort(tt.args.sparse); !reflect.DeepEqual(gotDense, tt.wantDense) {
				t.Errorf("sToDenseShort() = %v, want %v", gotDense, tt.wantDense)
			}
		})
	}
}
