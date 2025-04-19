package btreeReverse

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		result = append(result, node.Val)
	}

	dfs(root)
	return result
}

/*
*
Дан корень бинарного дерева. Нужно вернуть значения узлов в обратном порядке

Алгоритм обратного обхода:

1. Если текущий узел пустой - возвращаемся в родительскую вершину
2. Обходим левое поддерево рекурсивно
3. Обходим правое поддерево рекурсивно
4. Добавляем значение текущего узла в ответ
5. Возвращаемся в родительскую вершину
ВАЖНО: реализуй обход с использованием рекурсии
*/
func main() {
	/*
	         1
	        / \
	       2   3
	      / \   \
	     4   5   6

	   Обратный обход: [4 5 2 6 3 1]
	*/

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}

	result := postorderTraversal(root)
	fmt.Println(result) // ожидаемый результат: [4 5 2 6 3 1]
}
