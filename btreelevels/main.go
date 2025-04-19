package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(result) == level {
			result = append(result, []int{})
		}
		result[level] = append(result[level], node.Val)
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}

	dfs(root, 0)
	return result
}

/*
*
Дан корень бинарного дерева. Нужно вернуть значения всех узлов по уровням (слева-направо и сверху-вниз)

ВАЖНО: реши задачу с использованием рекурсии
*/
func main() {
	/*
	        1
	       / \
	      2   3
	     / \    \
	    4   5    6
	       / \
	      7   8

	   Ожидаемый вывод: [[1], [2, 3], [4, 5, 6], [7, 8]]
	*/

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 8}

	result := levelOrder(root)
	fmt.Println(result) // [[1], [2 3], [4 5 6], [7 8]]
}
