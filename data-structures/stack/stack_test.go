package stack

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	tests := []struct {
		name          string
		startingStack []int
		newVal        int
		want          []int
	}{
		{
			"Simple Test",
			[]int{1, 2, 3},
			4,
			[]int{1, 2, 3, 4},
		},
		{
			"Empty Stack",
			[]int{},
			2,
			[]int{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack{
				slice: tt.startingStack,
			}

			s.Push(tt.newVal)

			if !reflect.DeepEqual(s.slice, tt.want) {
				t.Errorf("got: %v, want: %v", s.slice, tt.want)
			}
		})
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		name          string
		startingStack []int
		wantVal       int
		wantStack     []int
		err           bool
	}{
		{
			"Simple Test",
			[]int{1, 2, 3, 4},
			4,
			[]int{1, 2, 3},
			false,
		},
		{
			"Single Item Stack",
			[]int{2},
			2,
			[]int{},
			false,
		},
		{
			"Empty Stack, Should Error",
			[]int{},
			0,
			[]int{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack{
				slice: tt.startingStack,
			}

			val, err := s.Pop()

			// Testing If Error Occurred
			if err != nil {
				if tt.err {
					return
				} else {
					t.Fatalf("Unexpected Error: %v", err.Error())
				}
			}

			// Checking That Value Popped is the One we Want
			if val != tt.wantVal {
				t.Errorf("got: %v, want: %v", val, tt.wantVal)
			}

			// Checking That the Value was Actually Popped From the Stack
			if !reflect.DeepEqual(s.slice, tt.wantStack) {
				t.Errorf("got: %v, want: %v", s.slice, tt.wantStack)
			}
		})
	}
}

func TestPeek(t *testing.T) {
	tests := []struct {
		name          string
		startingStack []int
		wantVal       int
		wantStack     []int
		err           bool
	}{
		{
			"Simple Test",
			[]int{1, 2, 3, 4},
			4,
			[]int{1, 2, 3, 4},
			false,
		},
		{
			"Single Item Stack",
			[]int{2},
			2,
			[]int{2},
			false,
		},
		{
			"Empty Stack, Should Error",
			[]int{},
			0,
			[]int{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack{
				slice: tt.startingStack,
			}

			val, err := s.Peek()

			// Testing If Error Occurred
			if err != nil {
				if tt.err {
					return
				} else {
					t.Fatalf("Unexpected Error: %v", err.Error())
				}
			}

			// Checking That Value Popped is the One we Want
			if val != tt.wantVal {
				t.Errorf("got: %v, want: %v", val, tt.wantVal)
			}

			// Checking That the Value was Actually Popped From the Stack
			if !reflect.DeepEqual(s.slice, tt.wantStack) {
				t.Errorf("got: %v, want: %v", s.slice, tt.wantStack)
			}
		})
	}
}
