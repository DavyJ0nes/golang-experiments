package main

import (
	"fmt"
	"runtime"
)

func main() {
	input := "Golang Experiments"
	fmt.Printf("Input:\t%s\n\n", input)
	fmt.Println("Num of GoRoutines", runtime.NumGoroutine())

	output := Reverse(input)
	fmt.Println("\nOutput:\t", output)
}

// Reverse takes a string and reverses it
// Has a bunch of annotations to help visualise what the loop is doing
func Reverse(s string) string {
	chars := []rune(s)
	iterations := 0
	fmt.Printf("---- char length: %d\n", len(chars))

	// l == left side of slice, r == right side of slice
	// after each iteration move towards middle of slice
	for l, r := 0, len(chars)-1; l < r; l, r = l+1, r-1 {
		fmt.Printf("l: %v\t\t[%v]\t(%v) | r: %v\t\t[%v]\t(%v)\n", chars[l], l, string(chars[l]), chars[r], r, string(chars[r]))
		chars[l], chars[r] = chars[r], chars[l]
		iterations++
	}

	fmt.Println("---- iterations:", iterations)

	return string(chars)
}
