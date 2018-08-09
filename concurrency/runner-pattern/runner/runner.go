package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner executes a set of tasks within a given timeout and can be shutdown
// using an OS system interrupt
type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	tasks     []func(int)
}

var (
	// ErrTimeout is returned when a value is received on the timeout channel
	ErrTimeout = errors.New("received timeout")
	// ErrInterrupt is returned when an event is received from the OS
	ErrInterrupt = errors.New("received interrupt")
)

// New returns a new runner that's ready to go
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add attaches tasks to a Runner
// A task is a function that takes an int ID
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start runs all the tasks and monitoring channel events
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

// run executes each of the registered tasks
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		task(id)
	}
	return nil
}

// gotInterrupt verifies whether an interrupt signal has been received from the OS
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
