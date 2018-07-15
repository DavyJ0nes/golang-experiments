package main

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		l []int
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"basic test",
			args{[]int{1, 3, 5, 6, 7, 8}, 5},
			2,
		},
		{
			"unknown item test",
			args{[]int{1, 3, 5, 6, 7, 8}, 12},
			-1,
		},
		{
			"long list test",
			args{
				[]int{1, 3, 5, 6, 7, 8, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
				12,
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.l, tt.args.i); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
