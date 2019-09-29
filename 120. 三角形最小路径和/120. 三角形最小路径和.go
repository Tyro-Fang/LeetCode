package main

import (
	"fmt"
)

func main() {
	a := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	b := minimumTotal(a)

	fmt.Println(b)
}
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			continue
		}
		for j := 0; j < i+1; j++ {
			triangle[i][j] = min(triangle[i+1][j], triangle[i+1][j+1]) + triangle[i][j]
		}
	}
	return triangle[0][0]
}
func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
