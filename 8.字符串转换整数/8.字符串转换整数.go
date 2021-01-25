package main

import (
	"fmt"
	"math"
)

func main() {
	//a := []int{1, 1, 1, 0}
	//nextPermutation(a)
	a := '4'
	b := a - '0'
	//b := myAtoi(a)
	fmt.Println(b)
	//c := maximalRectangle(a)
	//fmt.Println(c)
}
func myAtoi(str string) int {
	var res int
	start := false
	flag := 1
	flagNum := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			if start || flagNum > 0 {
				break
			}
			continue
		} else if str[i] == '+' {
			if start || flagNum > 0 {
				break
			}
			flagNum++
			continue
		} else if str[i] == '-' {
			if start || flagNum > 0 {
				break
			}
			flag = -1
			flagNum++
			continue
		} else if str[i] <= '9' && str[i] >= '0' {
			start = true
			val := (int)(str[i]-'0') * flag
			if res > math.MaxInt32/10 || (res == math.MaxInt32/10 && val > math.MaxInt32%10) {
				return math.MaxInt32
			} else if res < math.MinInt32/10 || (res == math.MinInt32/10 && val < math.MinInt32%10) {
				return math.MinInt32
			} else {
				res = (res*10 + val)
			}
		} else {
			break
		}
	}
	return res
}
