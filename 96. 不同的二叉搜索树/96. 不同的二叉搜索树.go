package main

import (
	"fmt"
)

func main() {
	a := 3
	b := numTrees(a)
	fmt.Println(b)

}

func numTrees(n int) int {
	nums := make([]int, n+1)
	nums[0], nums[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i+1; j++ {
			nums[i] += nums[j-1] * nums[i-j]
		}
	}
	return nums[n]
}
