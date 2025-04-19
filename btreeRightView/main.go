package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var result []int

	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if level == len(result) {
			result = append(result, node.Val)
		}
		dfs(node.Right, level+1)
		dfs(node.Left, level+1)
	}

	dfs(root, 0)
	return result
}

/*
*
Дан корень бинарного дерева. Нужно вернуть массив значений, где каждое значение соответствует самой правой вершине уровня дерева

ВАЖНО: реши задачу с использованием рекурсии
*/
func main() {
	/*
	         1
	        / \
	       2   3
	      / \
	     6   4
	    /
	   8

	    Уровни:
	    0 → 1
	    1 → 3
	    2 → 4
	    3 → 8

	    Вывод: [1, 3, 4, 8]
	*/

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 4}
	root.Left.Left.Left = &TreeNode{Val: 8}

	result := rightSideView(root)
	fmt.Println(result) // ожидаемый результат: [1 3 4 8]
}
