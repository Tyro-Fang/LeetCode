package main

import (
	"fmt"
)

func main() {
	a := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	b := maximalRectangle(a)
	fmt.Println(b)
	//c := maximalRectangle(a)
	//fmt.Println(c)
}
func maximalRectangle(matrix [][]byte) int {
	var width int = len(matrix[0])
	var height int = len(matrix)
	maxW := make([][]int, height)
	for i := 0; i < height; i++ {
		maxH := make([]int, width)
		for j := 0; j < width; j++ {
			if matrix[i][j] == '0' {
				maxH[j] = 0
			} else if j == 0 {
				maxH[j] = 1
			} else {
				maxH[j] = maxH[j-1] + 1
			}
		}
		maxW[i] = maxH
	}
	area := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			sum := 0
			minW := maxW[i][j]
			for k := i; k > -1; k-- {
				if minW == 0 {
					break
				}
				minW = minVal(maxW[k][j], minW)
				sum = (i - k + 1) * minW
				area = maxVal(area, sum)
			}
		}
	}
	return area
}
func maxVal(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
func minVal(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
