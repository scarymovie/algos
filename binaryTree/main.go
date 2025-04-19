package binaryTree

import "fmt"

/**
Дан корень бинарного дерева. Нужно вернуть значения узлов в прямом порядке


Алгоритм прямого обхода:

1. Если текущий узел пустой - возвращаемся в родительскую вершину
2. Добавляем значение текущего узла в ответ
3. Обходим левое поддерево рекурсивно
4. Обходим правое поддерево рекурсивно
5. Возвращаемся в родительскую вершину
*/

func main() {
	/*
	         1
	        / \
	       2   6
	      / \    \
	     3   5    7
	    /
	   4
	*/

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}
	root.Left.Left.Left = &TreeNode{Val: 4}

	result := preorderTraversal(root)
	fmt.Println(result) // ожидаемый результат: [1 2 3 4 5 6 7]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var result []int

	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)

		traversal(node.Right)

		result = append(result, node.Val)
	}
	traversal(root)

	return result
}
