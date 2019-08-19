package main

import (
	"fmt"
)

func main() {
	a := []int{3, 1, 5, 8}
	b := maxCoins(a)
	fmt.Println(b)
}
func maxCoins(nums []int) int {

	nums = append([]int{1}, nums...)
	nums = append(nums, []int{1}...)
	nlen := len(nums)
	res := make([][]int, nlen)
	for i := range res {
		res[i] = make([]int, nlen)
	}
	for l := 2; l < nlen; l++ {
		for i := 0; i < nlen-l; i++ {
			j := i + l
			for k := i + 1; k < j; k++ {
				res[i][j] = max(res[i][j], nums[i]*nums[k]*nums[j]+res[i][k]+res[k][j])
			}
		}

	}
	return res[0][nlen-1]
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
