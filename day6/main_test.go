package main

import (
	"reflect"
	"testing"
)

func Test_balance(t *testing.T) {
	type args struct {
		vs []int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int
	}{
		{"example1", args{[]int{0, 2, 7, 0}}, []int{2, 4, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes := balance(tt.args.vs)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("balance() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_getMaxIdx(t *testing.T) {
	type args struct {
		vs []int
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
	}{
		{"example1", args{[]int{0, 2, 7, 0}}, 2},
		{"4 to 0", args{[]int{4, 3, 2, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMax := getMaxIdx(tt.args.vs)
			if gotMax != tt.wantMax {
				t.Errorf("getMinMaxIdx() gotMax = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func Test_hasCycle(t *testing.T) {
	type args struct {
		vs []int
	}
	tests := []struct {
		name       string
		args       args
		wantCount  int
		wantCycles int
	}{
		{"example1", args{[]int{0, 2, 7, 0}}, 5, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCount, gotCycles := hasCycle(tt.args.vs)
			if gotCount != tt.wantCount {
				t.Errorf("hasCycle() gotCount = %v, want %v", gotCount, tt.wantCount)
			}
			if gotCycles != tt.wantCycles {
				t.Errorf("hasCycle() gotCycles = %v, want %v", gotCycles, tt.wantCycles)
			}
		})
	}
}
