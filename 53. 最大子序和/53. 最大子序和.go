package main

import (
	"fmt"
)

func main() {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	c := maxSubArrayFZ(a)
	fmt.Println(c)
}
func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0
	for _, val := range nums {
		if sum <= 0 {
			sum = val
		} else {
			sum += val
		}
		max = maxVal(max, sum)
	}
	return max
}
func maxVal(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxSubArrayFZ(nums []int) int {
	nlen := len(nums)
	if nlen == 1 {
		return nums[0]
	}
	var h int = nlen / 2
	maxleft := maxSubArrayFZ(nums[0:h])
	maxright := maxSubArrayFZ(nums[h:nlen])
	max1 := nums[h-1]
	sum1 := 0
	for i := h - 1; i >= 0; i-- {
		sum1 += nums[i]
		max1 = maxVal(max1, sum1)
	}
	max2 := nums[h]
	sum2 := 0
	for i := h; i < nlen; i++ {
		sum2 += nums[i]
		max2 = maxVal(max2, sum2)
	}
	max1 = max1 + max2
	max1 = maxVal(max1, max2)
	max1 = maxVal(max1, maxleft)
	max1 = maxVal(max1, maxright)
	return max1
}
