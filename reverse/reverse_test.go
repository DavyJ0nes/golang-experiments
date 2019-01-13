package main

import (
	"runtime"
	"testing"
)

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"palindrome",
			args{"racecar"},
			"racecar",
		},
		{
			"empty",
			args{""},
			"",
		},
		{
			"non palindrom",
			args{"golang"},
			"gnalog",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.s); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
			t.Log("Num of GoRoutines", runtime.NumGoroutine())
		})
	}
}

func BenchmarkReverseGolang(b *testing.B) {
	// run the Reverse function b.N times
	for n := 0; n < b.N; n++ {
		Reverse("Golang")
	}
}
