package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{3, 5, 2, 1, 6, 4}
	b := wiggleSortSwap1(a)
	//b := myAtoi(a)
	fmt.Println(b)

}
func wiggleSort(nums []int) []int {
	sort.Ints(nums)
	nlen := len(nums)
	if nlen <= 1 {
		return nums
	}
	i := 1
	for i < nlen {
		if i+1 >= nlen {
			break
		}
		nums[i], nums[i+1] = nums[i+1], nums[i]
		i += 2
	}
	return nums
}
func wiggleSortSwap(nums []int) []int {
	nlen := len(nums)
	if nlen <= 1 {
		return nums
	}
	i := 1
	for i < nlen {
		if nums[i-1] > nums[i] {
			nums[i-1], nums[i] = nums[i], nums[i-1]
		}
		if i+1 < nlen && nums[i+1] > nums[i] {
			nums[i+1], nums[i] = nums[i], nums[i+1]
		}
		i += 2
	}
	return nums
}

func wiggleSortSwap1(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		if i&1 == 1 {
			if nums[i] < nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		} else if i&1 == 0 {
			if nums[i] > nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
	}
	return nums
}
