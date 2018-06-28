package main

import "testing"

func TestTree(t *testing.T) {
	// Am just testing the error at the moment but will need to change this in future
	tests := []struct {
		name string
		arg  string
		want error
	}{
		{
			"test with local directory",
			".",
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tree(tt.arg); got != tt.want {
				t.Errorf("tree() = %v, want %v", got, tt.want)
			}
		})
	}
}
