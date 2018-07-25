package recursion

// Simple Looping
func shapeAreaLooping(n int) int {
	area := 1
	for i := 1; i <= n; i++ {
		area += (i * 4) - 4
	}

	return area
}

// With Recursion
func shapeAreaRecursion(n int) int {
	if n == 1 {
		return 1
	} else {
		return ((n * 4) - 4) + shapeAreaRecursion(n-1)
	}
}

// add all elements of a slice together using recursion
func sumSlice(input []int) int {
	var x, sum int

	if len(input) == 0 {
		return 0
	}

	x, input = input[0], input[1:]
	sum = x + sumSlice(input)

	return sum
}

// count elements in slice using recursion
func sliceCount(input []string) int {
	var count int

	if len(input) == 0 {
		return 0
	}

	input = input[1:]
	count = sliceCount(input) + 1

	return count
}
