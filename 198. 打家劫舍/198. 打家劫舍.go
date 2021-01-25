package main

import "fmt"

func main() {
	a := []int{2, 1, 1, 2}
	b := rob(a)
	//c := maximalRectangle(a)
	fmt.Println(b)
}

func rob(nums []int) int {
	rlen := len(nums)
	if rlen == 0 {
		return 0
	} else if rlen == 1 {
		return nums[0]
	} else if rlen == 2 {
		return max(nums[0], nums[1])
	}
	for i := 2; i < rlen; i++ {
		if i == 2 {
			nums[i] = nums[i-2] + nums[i]
		} else {
			nums[i] = max(nums[i-2], nums[i-3]) + nums[i]
		}
	}
	return max(nums[rlen-1], nums[rlen-2])
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
