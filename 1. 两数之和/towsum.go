package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := 5
	c := twoSum(a, b)
	fmt.Println(c)
}
func twoSum(nums []int, target int) []int {
	var m = make(map[int]int)
	var res = make([]int, 0)
	for key, val := range nums {
		t := target - val
		_, ok := m[t]
		if ok {
			res = append(res, m[t])
			res = append(res, key)
			return res
		}
		m[val] = key
	}
	return res
}
