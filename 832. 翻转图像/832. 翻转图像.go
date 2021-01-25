package main

import "fmt"

func main() {
	a := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	b := flipAndInvertImage(a)
	fmt.Println(b)
	//c := maximalRectangle(a)
	//fmt.Println(c)
}
func flipAndInvertImage(A [][]int) [][]int {
	xlen := len(A)
	if xlen < 1 {
		return A
	}
	ylen := len(A[0])
	if ylen < 1 {
		return A
	}
	for i := 0; i < xlen; i++ {
		for j, k := 0, ylen-1; j <= k; {
			A[i][j], A[i][k] = 1-A[i][k], 1-A[i][j]
			j++
			k--
		}
	}
	return A
}
