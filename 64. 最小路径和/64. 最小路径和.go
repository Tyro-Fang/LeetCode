package main

import (
	"fmt"
)

func main() {
	a := [][]int{{1, 2, 5}, {3, 2, 1}}
	b := minPathSum(a)

	fmt.Println(b)
}
func minPathSum(grid [][]int) int {
	n := len(grid)
	m := 0
	if n > 0 {
		m = len(grid[0])
	}
	for j := n - 1; j >= 0; j-- {
		for i := m - 1; i >= 0; i-- {
			if j == n-1 && i == m-1 {
				continue
			}
			if j == n-1 {
				grid[j][i] = grid[j][i] + grid[j][i+1]
				continue
			}
			if i == m-1 {
				grid[j][i] = grid[j][i] + grid[j+1][i]
				continue
			}
			grid[j][i] = min(grid[j][i+1], grid[j+1][i]) + grid[j][i]
		}
	}
	return grid[0][0]
}
func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
