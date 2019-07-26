package main

import (
	"fmt"
)

func main() {
	a := []int{2, 3, 1}
	b := jump(a)
	//b := myAtoi(a)
	fmt.Println(b)
}
func jump(nums []int) int {
	i := 0
	res := 0
	nlen := len(nums)
	if nlen <= 1 {
		return 0
	}
	for i < nlen {
		pos := 0
		step := 0
		if i == nlen-1 {
			break
		}
		if i+nums[i] >= nlen {
			res++
			break
		}
		for j := 1; j < nums[i]+1; j++ {
			if j < nlen && j > step {
				step = nums[i+j]
				pos = i + j
			}
		}
		i = pos
		res++
	}
	return res
}
