package main

import (
	"fmt"
)

//TreeNode is a treenode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := 5
	b := generateTrees(a)
	fmt.Println(b)

}

func generateTrees(n int) []*TreeNode {
	nums := []TreeNode{}
	for i := 1; i <= n; i++ {
		temp := TreeNode{
			Val:   i,
			Left:  nil,
			Right: nil,
		}
		nums = append(nums, temp)
	}
	b := generateTree(nums)
}
func generateTree(nums []TreeNode) []*TreeNode {
	nlen := len(nums)
	if nlen == 1 {
		return []*TreeNode{&nums[0]}
	}

}
