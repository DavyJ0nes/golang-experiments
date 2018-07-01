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
