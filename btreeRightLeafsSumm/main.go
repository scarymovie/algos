package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfRightLeaves(root *TreeNode) int {
	var sum int

	var dfs func(node *TreeNode, isRight bool)
	dfs = func(node *TreeNode, isRight bool) {
		if node == nil {
			return
		}
		if isRight {
			if node.Left == nil && node.Right == nil {
				sum += node.Val
			}
		}
		dfs(node.Right, true)
		dfs(node.Left, false)
	}

	dfs(root, false)
	return sum
}

func main() {
	/*
	        1
	       / \
	      2   3
	     / \    \
	    4   5    6
	       / \
	      7   8

	   Правые листья: 8, 6 → сумма: 14
	*/

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 8}

	result := sumOfRightLeaves(root)
	fmt.Println(result) // ожидаемый результат: 14
}
