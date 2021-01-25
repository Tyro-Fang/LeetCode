package main

import "fmt"

func main() {

	wordDict := []int{-2, 0, -1}
	b := permute(wordDict)
	//c := maximalRectangle(a)
	fmt.Println(b)
}
func permute(nums []int) [][]int {
	nlen := len(nums)
	if nlen <= 1 {
		return [][]int{nums}
	}
	return pailie(nums)
}
func pailie(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	var res [][]int
	s := pailie(nums[1:])
	for _, v := range s {
		for j := 1; j <= len(v); j++ {
			a := make([]int, len(v[0:j]))
			b := make([]int, len(v[j:]))
			copy(a, v[0:j])
			copy(b, v[j:])
			temp := append(a, nums[0])
			temp = append(temp, b...)
			res = append(res, temp)
		}
		temp := []int{nums[0]}
		res = append(res, append(temp, v...))
	}
	return res
}
