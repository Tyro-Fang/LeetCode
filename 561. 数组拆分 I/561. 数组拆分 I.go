package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{1, 3, 4, 2}
	b := arrayPairSum(a)
	//b := myAtoi(a)
	fmt.Println(b)

}
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	ans := 0
	i := 0
	for i < len(nums) {
		ans += nums[i]
		i += 2
	}
	return ans
}
