package recursion

import "testing"

func TestLoopingArea(t *testing.T) {
	tests := []struct {
    name string
		input int
		want int
	}{
		{"1",1,1},
		{"2",2,5},
		{"3",3,13},
		{"4",4,25},
		{"5",5,41},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shapeAreaLooping(tt.input); got != tt.want {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}
