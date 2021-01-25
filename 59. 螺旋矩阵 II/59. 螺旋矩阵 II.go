package main

import "fmt"

func main() {
	a := 3

	b := generateMatrix(a)
	fmt.Println(b)
	//c := maximalRectangle(a)
	//fmt.Println(c)
}
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	l := (n + 1) / 2 //执行次数
	count := 1
	for i := 0; i < l; i++ {
		for j := i; j < n-i; j++ {
			res[i][j] = count
			count++
		}
		for j := i + 1; j < n-i; j++ {
			res[j][n-i-1] = count
			count++
		}
		for j := n - i - 2; j >= i; j-- {
			res[n-i-1][j] = count
			count++
		}
		for j := n - i - 2; j >= i+1; j-- {
			res[j][i] = count
			count++
		}
	}
	return res
}
