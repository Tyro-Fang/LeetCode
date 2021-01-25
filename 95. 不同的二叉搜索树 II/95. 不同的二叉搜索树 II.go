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
	if n == 0 {
		return make([]*TreeNode, 0)
	}
	return generateTree(1, n)
}
func generateTree(start int, end int) []*TreeNode {
	trees := make([]*TreeNode, 0)
	if start > end {
		trees = append(trees, nil)
		return trees
	}
	for i := start; i <= end; i++ {
		left := generateTree(start, i-1)
		right := generateTree(i+1, end)
		for _, j := range left {
			for _, k := range right {
				root := TreeNode{i, nil, nil}
				root.Left = j
				root.Right = k
				trees = append(trees, &root)
			}
		}
	}
	return trees

}
