package main

import (
	"fmt"
	"sort"
)

var ans [][]int

func main() {
	a := []int{2, 3, 5}
	c := 8
	b := combinationSum(a, c)
	fmt.Println(b)
	//c := maximalRectangle(a)
	//fmt.Println(c)

}
func combinationSum(candidates []int, target int) [][]int {
	sort.Sort(sort.Reverse(sort.IntSlice(candidates)))
	ans = [][]int{}
	cw := 0
	cur := []int{}
	dfs(candidates, target, cur, cw)
	return ans
}

func dfs(candidates []int, target int, cur []int, cw int) {
	if cw == target {
		clen := len(cur)
		temp := make([]int, clen)
		copy(temp, cur)
		ans = append(ans, temp)
		return
	} else if cw > target {
		return
	} else {
		clen := len(candidates)
		for i := 0; i < clen; i++ {
			t := len(cur)
			if t == 0 || candidates[i] <= cur[t-1] {
				cur = append(cur, candidates[i])
				cw += candidates[i]
				dfs(candidates, target, cur, cw)
				cw -= candidates[i]
				cur = cur[0 : len(cur)-1]
			}
		}
	}
}
