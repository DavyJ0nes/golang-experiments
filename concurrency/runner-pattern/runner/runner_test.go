package runner

import (
	"testing"
	"time"
)

func TestBasicRunner(t *testing.T) {
	// Create New Runner
	runner := New(time.Duration(1 * time.Second))

	// Add simple function to runner
	runner.Add(func(i int) {
		i = i * 2
	})

	want := 1
	got := len(runner.tasks)

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}

	// Start task runner and check error
	err := runner.Start()
	if err != nil {
		t.Errorf("Unexpected Error: %v", err)
	}

}
