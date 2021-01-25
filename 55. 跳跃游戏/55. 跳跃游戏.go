package main

import (
	"fmt"
)

func main() {
	wordDict := []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0}
	b := canJump(wordDict)
	fmt.Println(b)
}
func canJump(nums []int) bool {
	nlen := len(nums)
	if nlen <= 1 {
		return true
	}
	i := 0
	for {
		step := nums[i]
		if step == 0 {
			return false
		}
		max := 0
		maxI := 0
		if step+i >= nlen-1 {
			return true
		}
		for j := 1; j <= step; j++ {
			if nums[i+j]+j > max {
				max = nums[i+j] + j
				maxI = i + j
			}
		}
		if i+max >= nlen {
			return true
		}
		i = maxI
	}
	return false
}
