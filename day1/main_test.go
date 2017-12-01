package main

import (
	"testing"
)

func Test_antiCaptcha(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{"1122"}, 3},
		{"example2", args{"1111"}, 4},
		{"example3", args{"1234"}, 0},
		{"example4", args{"91212129"}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := antiCaptcha(tt.args.in); got != tt.want {
				t.Errorf("antiCaptcha() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_antiCaptchaHalfway(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{"1212"}, 6},
		{"example2", args{"1221"}, 0},
		{"example3", args{"123123"}, 12},
		{"example4", args{"12131415"}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := antiCaptchaHalfway(tt.args.in); got != tt.want {
				t.Errorf("antiCaptchaHalfway() = %v, want %v", got, tt.want)
			}
		})
	}
}
