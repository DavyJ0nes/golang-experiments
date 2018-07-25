package recursion

import "testing"

func TestLoopingArea(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"1", 1, 1},
		{"2", 2, 5},
		{"3", 3, 13},
		{"4", 4, 25},
		{"5", 5, 41},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shapeAreaLooping(tt.input); got != tt.want {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}

func TestRecursiveArea(t *testing.T) {
	tests := []struct {
		name  string
		input int
		want  int
	}{
		{"1", 1, 1},
		{"2", 2, 5},
		{"3", 3, 13},
		{"4", 4, 25},
		{"5", 5, 41},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shapeAreaRecursion(tt.input); got != tt.want {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}

func Test_sumSlice(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"basic test",
			args{input: []int{2, 4, 6}},
			12,
		},
		{
			"empty input test",
			args{input: []int{}},
			0,
		},
		{
			"longer input test",
			args{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
			45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSlice(tt.args.input); got != tt.want {
				t.Errorf("sumSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceCount(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"basic test",
			args{input: []string{"do this", "and that"}},
			2,
		},
		{
			"longer test",
			args{input: []string{"a", "b", "c", "d", "e", "f", "g"}},
			7,
		},
		{
			"empty input test",
			args{input: []string{}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceCount(tt.args.input); got != tt.want {
				t.Errorf("sliceCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
