package main

import "fmt"

func main() {
	a := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	b := uniquePathsWithObstacles(a)
	fmt.Println(b)
}
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}
	n := len(obstacleGrid[0])
	if m == 0 || n == 0 {
		return 0
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
				continue
			}
			if i == 0 && j == 0 {
				if obstacleGrid[0][0] == 1 {
					return 0
				}
				obstacleGrid[0][0] = 1
				continue
			}
			if i == 0 {
				obstacleGrid[i][j] = obstacleGrid[i][j-1]
				continue
			}
			if j == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][j]
				continue
			}
			obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
		}
	}
	return obstacleGrid[m-1][n-1]
}
