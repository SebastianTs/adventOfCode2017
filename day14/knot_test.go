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
func Test_sToDenseBlock(t *testing.T) {
	type args struct {
		sparse []byte
	}
	tests := []struct {
		name      string
		args      args
		wantDense byte
	}{
		{"ExampleBlock", args{[]byte{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22}}, byte(64)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDense := sToDenseBlock(tt.args.sparse); !reflect.DeepEqual(gotDense, tt.wantDense) {
				t.Errorf("sToDenseShort() = %v, want %v", gotDense, tt.wantDense)
			}
		})
	}
}

func Test_knotString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"emptyString", args{""}, "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC2017String", args{"AoC 2017"}, "33efeb34ea91902bb2f59c9920caa6cd"},
		{"OneToThree", args{"1,2,3"}, "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"OneToFour", args{"1,2,4"}, "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := knotString(tt.args.s); got != tt.want {
				t.Errorf("knotString() = %v, want %v", got, tt.want)
			}
		})
	}
}
