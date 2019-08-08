package main

import (
	"fmt"
)

func main() {

	b := minDistance("horse", "ros")
	fmt.Println(b)

}
func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)
	w1, w2 := []byte(word1), []byte(word2)
	res := make([][]int, len1+1)
	for i := range res {
		res[i] = make([]int, len2+1)
	}
	left, up, leftAndUp, val := 0, 0, 0, 0
	for i := 0; i < len1+1; i++ {
		for j := 0; j < len2+1; j++ {
			if i == 0 {
				res[i][j] = j
				continue
			}
			if j == 0 {
				res[i][j] = i
				continue
			}
			left = res[i][j-1]
			up = res[i-1][j]
			leftAndUp = res[i-1][j-1]
			if w1[i-1] == w2[j-1] {
				leftAndUp--
			}
			val = min(left, up, leftAndUp)
			res[i][j] = val + 1
		}
	}
	return res[len1][len2]
}
func min(a int, b int, c int) int {
	temp := a
	if temp > b {
		temp = b
	}
	if temp > c {
		temp = c
	}
	return temp
}
