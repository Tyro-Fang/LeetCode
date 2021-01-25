package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{7, 1, 5, 3, 6, 4}
	b := maxProfit(a)

	fmt.Println(b)
}
func maxProfit(prices []int) int {
	res := 0
	nlen := len(prices)
	if nlen <= 1 {
		return 0
	}
	min := math.MaxInt32
	for i := 0; i < nlen; i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}
		if prices[i]-min > res {
			res = prices[i] - min
		}
	}
	return res
}
