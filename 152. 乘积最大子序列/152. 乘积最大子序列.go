package main

import "fmt"

func main() {

	wordDict := []int{-2, 0, -1}
	b := maxProduct(wordDict)
	//c := maximalRectangle(a)
	fmt.Println(b)
}
func maxProduct(nums []int) int {
	nlen := len(nums)
	max := nums[0]
	min := nums[0]
	res := nums[0]
	for i := 1; i < nlen; i++ {
		temp1 := maxInt(maxInt(max*nums[i], nums[i]), min*nums[i])
		temp2 := minInt(minInt(min*nums[i], nums[i]), max*nums[i])
		max = maxInt(temp1, temp2)
		min = minInt(temp1, temp2)
		if res < max {
			res = max
		}
	}
	return res
}
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minInt(a, b int) int {
	if a > b {
		return b
	}
	return a
}
