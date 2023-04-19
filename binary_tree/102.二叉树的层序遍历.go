package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 广度优先搜索
func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	var (
		current = root
		next    = make([]*TreeNode, 0)
	)

	ans := make([][]int, 0)

	depth := 0

	next = append(next, current)

	for len(next) > 0 {
		ans = append(ans, make([]int, 0))

		tmp := make([]*TreeNode, 0)

		for _, node := range next {
			ans[depth] = append(ans[depth], node.Val)
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}

			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}

		depth++

		next = tmp
	}

	return ans
}

// 广度优先搜索
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	q := Queue{
		root,
	}

	ans := make([][]int, 0)

	for !q.IsEmpty() {
		length := q.Length()
		ans = append(ans, []int{})
		for i := 0; i < length; i++ {
			tmp := q.Pop()
			node := tmp.(*TreeNode)
			ans[len(ans)-1] = append(ans[len(ans)-1], node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
		}
	}

	return ans
}
