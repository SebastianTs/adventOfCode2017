package main

import "testing"

func Test_spinlock(t *testing.T) {
	type args struct {
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{3}, 638},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spinlock(tt.args.steps); got != tt.want {
				t.Errorf("spinlock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_spinlock5E7(t *testing.T) {
	type args struct {
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"ThreeSteps", args{3}, 1222153},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spinlock5E7(tt.args.steps); got != tt.want {
				t.Errorf("spinlock5E7() = %v, want %v", got, tt.want)
			}
		})
	}
}
