package main

import (
	"testing"
)

func Test_isValid(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example1", args{[]string{"aa", "bb", "cc", "dd", "ee"}}, true},
		{"example2", args{[]string{"aa", "bb", "cc", "dd", "aa"}}, false},
		{"example3", args{[]string{"aa", "bb", "cc", "dd", "aaa"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidWithoutAnagram(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example1", args{[]string{"abcde", "fghij"}}, true},
		{"example2", args{[]string{"abcde", "xyz", "ecdab"}}, false},
		{"example3", args{[]string{"a", "ab", "abc", "abd", "abf", "abj"}}, true},
		{"example4", args{[]string{"iiii", "oiii", "ooii", "oooi", "oooo"}}, true},
		{"example5", args{[]string{"oiii", "ioii", "iioi", "iiio"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidWithoutAnagram(tt.args.s); got != tt.want {
				t.Errorf("isValidWithoutAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
