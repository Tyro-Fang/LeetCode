package main

import (
	"fmt"
)

func main() {
	wordDict := [][]int{{1, 4}, {2, 3}}
	b := merge(wordDict)
	fmt.Println(b)
}
func merge(intervals [][]int) [][]int {
	ilen := len(intervals)
	for i := ilen - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			ok, res := judge(intervals[i], intervals[j])
			if ok {
				intervals[j] = res
				intervals = append(intervals[:i], intervals[i+1:]...)
				break
			}
		}
	}
	return intervals
}
func judge(i1 []int, i2 []int) (bool, []int) {
	if (i2[0] >= i1[0] && i2[0] <= i1[1]) || (i2[1] >= i1[0] && i2[1] <= i1[1]) || (i1[1] >= i2[0] && i1[1] <= i2[1]) || (i1[0] >= i2[0] && i1[0] <= i2[1]) {
		var t1, t2 int
		if i1[0] < i2[0] {
			t1 = i1[0]
		} else {
			t1 = i2[0]
		}
		if i1[1] < i2[1] {
			t2 = i2[1]
		} else {
			t2 = i1[1]
		}
		return true, []int{t1, t2}
	}
	return false, []int{}
}
