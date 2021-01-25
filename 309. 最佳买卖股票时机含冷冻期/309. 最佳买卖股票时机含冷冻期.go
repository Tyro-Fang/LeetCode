package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 0, 2}
	b := maxProfit(a)
	//c := maximalRectangle(a)
	fmt.Println(b)
}
func maxProfit(prices []int) int {
	sold, hold, rest := 0, 0, 0
	for i := range prices {
		preSold := sold
		sold = hold + prices[i]
		hold = max(hold, rest-prices[i])
		rest = max(rest, preSold)
	}
	return max(sold, rest)
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
