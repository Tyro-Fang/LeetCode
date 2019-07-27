package main

import (
	"fmt"
)

func main() {
	a := []int{4, 1, 1, 3, 1, 1, 1}
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
		if i == nlen-1 { //跳到最后一步
			break
		}
		if i+nums[i] >= nlen { //若可以跳到最后一个元素，则跳到最后一个元素
			res++
			i = nlen - 1
			break
		}
		pos := 0
		step := 0
		for j := 1; j < nums[i]+1; j++ {
			if (j + i) == nlen-1 {
				step = nums[i+j]
				pos = i + j
				break
			}
			if j < nlen && (nums[i+j]+j) > step {
				step = nums[i+j] + j
				pos = i + j
			}
		}
		i = pos
		res++
	}
	return res
}
