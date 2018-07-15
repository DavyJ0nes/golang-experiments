package main

import "fmt"

func main() {
	list := []int{1, 3, 5, 7, 9, 12}
	fmt.Println("--------------------")
	fmt.Println(binarySearch(list, 3))
	fmt.Println("--------------------")
	fmt.Println(binarySearch(list, -1))
	fmt.Println("--------------------")
}

func binarySearch(list []int, item int) int {
	fmt.Println("want:", item)

	for low, high, i := 0, len(list)-1, 1; low <= high; i++ {
		mid := (low + high) / 2
		guess := list[mid]

		fmt.Println("iteration:", i)
		fmt.Println("mid point:", mid)
		fmt.Println("low:", low)
		fmt.Println("high:", high)
		fmt.Println("guess:", guess)
		if guess == item {
			return mid
		}

		if guess < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
