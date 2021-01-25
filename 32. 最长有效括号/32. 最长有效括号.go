package main

import (
	"fmt"
)

func main() {
	a := "()(())"
	b := longestValidParenthesesStack(a)

	fmt.Println(b)
}
func longestValidParentheses(s string) int {
	slen := len(s)
	res := make([]int, slen)
	max := 0
	for i := 1; i < slen; i++ {
		if s[i] == ')' {
			if s[i-1] == ')' {
				if i-res[i-1]-1 >= 0 {
					if s[i-res[i-1]-1] == '(' {
						if i-res[i-1]-2 >= 0 {
							res[i] = res[i-1] + res[i-res[i-1]-2] + 2
						} else {
							res[i] = res[i-1] + 2
						}
					}
				}
			} else {
				if i > 1 {
					res[i] = res[i-2] + 2
				} else {
					res[i] = 2
				}
			}
		}
		if res[i] > max {
			max = res[i]
		}
	}
	return max
}

func longestValidParenthesesStack(s string) int {
	slen := len(s)
	sl := []int{-1}
	max := 0
	for i := 0; i < slen; i++ {
		if s[i] == '(' {
			sl = append(sl, i)
		}
		if s[i] == ')' {
			snum := len(sl)
			if snum > 0 {
				sl = sl[:snum-1]
				snum1 := len(sl)
				if snum1 > 0 {
					if i-sl[snum1-1] > max {
						max = i - sl[snum1-1]
					}
				} else {
					sl = append(sl, i)
				}
			}
		}
	}
	return max
}
