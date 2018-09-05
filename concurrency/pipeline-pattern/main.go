// Implements an example pipeline within Go.
// Corresponding blog post can be found [here](https://blog.golang.org/pipelines)

package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

// gen pushes input values to channel in async way
func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}

		close(out)
	}()

	return out
}

// square reads from a channel, squares the value
// and pushes squared value to output channel
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}

		close(out)
	}()

	return out
}

// merge works by fanning in numerous channels into one
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	wg.Add(len(cs))

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	for _, c := range cs {
		go output(c)
	}

	// block closing of channel until
	// all inputs have been merged
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	trace.Start(os.Stderr)
	in := gen(2, 3)

	c1 := square(in)
	c2 := square(in)

	for val := range merge(c1, c2) {
		fmt.Println(val)
	}
	trace.Stop()
}
