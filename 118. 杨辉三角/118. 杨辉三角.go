package main

import (
	"fmt"
)

func main() {
	a := 5
	b := generate(a)
	fmt.Println(b)
	//c := maximalRectangle(a)
	//fmt.Println(c)

}
func generate(numRows int) [][]int {
	var res = [][]int{}
	if numRows == 0 {
		return res
	}
	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}
	for i := 1; i < numRows; i++ {
		temp := []int{}
		for j := 0; j < i+1; j++ {
			if j == 0 {
				temp = append(temp, res[i-1][0])
			} else if j == i {
				temp = append(temp, res[i-1][j-1])
			} else {
				temp = append(temp, res[i-1][j-1]+res[i-1][j])
			}
		}
		tCopy := make([]int, i+1)
		copy(tCopy, temp)
		res = append(res, tCopy)
	}
	return res
}
