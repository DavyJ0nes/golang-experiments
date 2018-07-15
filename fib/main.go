package main

import "fmt"

func main() {
	fib := fib()
	for i := 0; i < 1000; i++ {
		fmt.Println(fib())
	}
}

func fib() func() int64 {
	a, b := int64(0), int64(1)
	return func() int64 {
		a, b = b, a+b
		return a
	}
}
