package main

import (
	"fmt"
	"strconv"
	"strings"
)

type fizzBuzz struct {
	output []string
}

func main() {
	limit := 100
	// Create 3 instances of fizzbuzz.
	// 1 for each of the implementation functions
	fb1 := fizzBuzz{}
	fb2 := fizzBuzz{}
	fb3 := fizzBuzz{}

	// Generate the output for each of the implementations
	fb1.output = basicFizzBuzz(limit)
	fb2.output = improvedFizzBuzz(limit)
	for out := range chanFizzBuzz(limit) {
		fb3.output = append(fb3.output, out)
	}

	// Print output. Not sure why this is useful.
	fmt.Printf("Basic:\n%s\n\n", fb1)
	fmt.Printf("Improved:\n%s\n\n", fb2)
	fmt.Printf("Channel:\n%s\n", fb3)
}

// super basic implementation
func basicFizzBuzz(limit int) []string {
	output := []string{}
	for i := 1; i <= limit; i++ {
		if i%3 == 0 && i%5 == 0 {
			output = append(output, "FizzBuzz")
		} else if i%3 == 0 {
			output = append(output, "Fizz")
		} else if i%5 == 0 {
			output = append(output, "Buzz")
		} else {
			output = append(output, strconv.Itoa(i))
		}
	}

	return output
}

// slightly better implementation. One less check
func improvedFizzBuzz(limit int) []string {
	var output []string

	for i := 1; i <= limit; i++ {
		result := ""

		if i%3 == 0 {
			result += "Fizz"
		}

		if i%5 == 0 {
			result += "Buzz"
		}

		if result != "" {
			output = append(output, result)
			continue
		}

		output = append(output, strconv.Itoa(i))
	}
	return output
}

// using channels for this one
func chanFizzBuzz(limit int) <-chan string {
	output := make(chan string, limit)

	go func() {
		for i := 1; i <= limit; i++ {
			result := ""
			if i%3 == 0 {
				result += "Fizz"
			}
			if i%5 == 0 {
				result += "Buzz"
			}
			if result == "" {
				result = strconv.Itoa(i)
			}
			output <- result
		}
		close(output)
	}()
	return output
}

// String is used to adhere to the Stringer interface for better output
func (fb fizzBuzz) String() string {
	return strings.Join(fb.output, "\n")
}
