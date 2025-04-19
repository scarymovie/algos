package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeaves(root *TreeNode) int {
	var sum int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			sum += node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return sum
}

/*
*
Дан корень бинарного дерева. Нужно вернуть сумму всех листьев в дереве

ВАЖНО: реши задачу с использованием рекурсии

Пример 1:

Ввод: root = [1,2,3,4,5,null,6,null,null,7,8]
Вывод: 25
Объяснение: В дереве 4 листа со значениями [4,7,8,6], а их сумма 25
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

	   Листья: 4, 7, 8, 6 → сумма: 25
	*/

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 8}

	result := sumOfLeaves(root)
	fmt.Println(result) // ожидаемый результат: 25
}
