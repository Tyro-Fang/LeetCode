package main

import "fmt"

func main() {

	wordDict := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(wordDict)
	//c := maximalRectangle(a)
	//fmt.Println(b)
}
func rotate(matrix [][]int) {
	nlen := len(matrix)
	for i := 0; i < nlen/2; i++ {
		for j := i; j < nlen-1-i; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[nlen-1-j][i]
			temp, matrix[j][nlen-1-i] = matrix[j][nlen-1-i], temp
			temp, matrix[nlen-1-i][nlen-1-j] = matrix[nlen-1-i][nlen-1-j], temp
			matrix[nlen-1-j][i] = temp
		}
	}
	fmt.Println(matrix)
}
