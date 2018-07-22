package main

import "fmt"

func main() {
	matrix := [][]int{
		[]int{1, 1, 1, 0},
		[]int{0, 5, 0, 1},
		[]int{2, 1, 3, 10},
	}
	op := matrixElementsSum(matrix)
	fmt.Println(op)

}

func matrixElementsSum(matrix [][]int) int {
	var total int

	for si, s := range matrix {
		for i, v := range s {
			if v == 0 {
				for j := si; j < len(matrix); j++ {
					matrix[j][i] = 0
				}
			} else {
				total += v
			}
		}
	}
	return total
}
