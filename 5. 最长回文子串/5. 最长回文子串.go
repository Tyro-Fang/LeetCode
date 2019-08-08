package main

import (
	"fmt"
)

func main() {
	a := "ccc"
	b := longestPalindrome(a)
	fmt.Println(b)
	//c := maximalRectangle(a)
	//fmt.Println(c)

}
func longestPalindrome(s string) string {
	slen := len(s)
	if slen <= 1 {
		return s
	}
	maxLen := 0 //最大长度
	maxBeg := 0 //最长字符串开始index
	maxEnd := 0 //最长字符串结束index
	//分为两种情况，一种是奇数，以某个字符串为中心做回文，一种是偶数，即左右对称
	for i := 1; i < slen; i++ {
		j := i - 1
		k := i + 1
		l:=i
		temp := 1
		for j >= 0 && k < slen {
			if s[j] == s[k] {
				temp += 2
				if temp > maxLen {
					maxLen = temp
					maxBeg = j
					maxEnd = k
				}
				j--
				k++
			} else {
				break
			}
		}
		j = i - 1
		k = i
		temp = 0
		for j >= 0 && k < slen {
			if s[j] == s[k] {
				temp += 2
				if temp > maxLen {
					maxLen = temp
					maxBeg = j
					maxEnd = k
				}
				j--
				k++
			} else {

				break
			}
		}
	}
	return s[maxBeg : maxEnd+1]
}
