package main

import (
	"fmt"
)

func main() {

	b := climbStairs(3)
	//b := myAtoi(a)
	fmt.Println(b)
}
func climbStairs(n int) int {
	a, b, res := 1, 1, 0
	if n <= 1 {
		return 1
	}
	for i := 1; i < n; i++ {
		res = a + b
		a = b
		b = res
	}
	return res
}
