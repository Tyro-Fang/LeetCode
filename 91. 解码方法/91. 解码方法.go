package main

import "fmt"

func main() {
	a := "221112221111222111222"
	b := numDecodings(a)
	//c := maximalRectangle(a)
	fmt.Println(b)
}

func numDecodings(s string) int {
	slen := len(s)
	if s[0] == '0' {
		return 0
	}
	pre := 1
	cur := 1
	for i := 1; i < slen; i++ {
		temp := cur
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				cur = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			cur = cur + pre
		}
		pre = temp
	}
	return cur
}
