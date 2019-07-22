package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{1, 2, 3}
	b := subsetBit(a)
	fmt.Println(b)
	//c := maximalRectangle(a)
	//fmt.Println(c)

}

func subsets(nums []int) [][]int {
	xlen := len(nums)
	if xlen == 0 {
		return [][]int{}
	}
	var res = [][]int{}
	i := 0
	k := 0
	j := xlen - 1
	for i <= j {
		for k <= j+1 {
			res = append(res, nums[i:k])
			k++
		}
		i++
		k = i + 1
	}
	return res
}
func subsetsDG(nums []int) [][]int { //递归方案，若是数组长度较大，则有可能出现栈溢出
	xlen := len(nums)
	if xlen <= 1 {
		return [][]int{nums}
	}
	res := subsetsDG(nums[0 : xlen-1])
	rlen := len(res)
	for i := 0; i < rlen; i++ {
		temp1 := make([]int, len(res[i])+1)
		copy(temp1, append(res[i], nums[xlen-1]))
		res = append(res, temp1)
	}
	res = append(res, []int{nums[xlen-1]})
	return res
}
func subsetBit(nums []int) [][]int {
	nlen := len(nums)
	count := int(math.Pow(2, float64(nlen)))
	res := [][]int{}
	for i := 0; i < count; i++ {
		pos := 0
		flag := i
		temp := []int{}
		for j := 0; j < nlen; j++ {
			if (flag & 1) == 1 {
				temp = append(temp, nums[j])
			}
			flag = flag >> 1
			pos++
		}
		res = append(res, temp)
	}
	return res
}
