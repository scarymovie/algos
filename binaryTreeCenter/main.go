package binaryTreeCenter

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
*
Дан корень бинарного дерева. Нужно вернуть значения узлов в центрированном порядке

Алгоритм центрированного обхода:

1. Если текущий узел пустой - возвращаемся в родительскую вершину
2. Обходим левое поддерево рекурсивно
3. Добавляем значение текущего узла в ответ
4. Обходим правое поддерево рекурсивно
5. Возвращаемся в родительскую вершину
ВАЖНО: реализуй обход с использованием рекурсии
*/
func main() {
	/*
	         5
	        / \
	       3   6
	      / \    \
	     2   4    7
	    /
	   1
	*/

	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 7}
	root.Left.Left.Left = &TreeNode{Val: 1}

	result := inorderTraversal(root)
	fmt.Println(result) // ожидаемый результат: [1 2 3 4 5 6 7]
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	traversal(root, &res)
	return res
}

func traversal(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	traversal(root.Left, res)
	*res = append(*res, root.Val)
	traversal(root.Right, res)
}
