package binary_tree

func zigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	var (
		next  = make([]*TreeNode, 0)
		ans   = make([][]int, 0)
		depth = 0
	)

	next = append(next, root)

	for len(next) > 0 {
		ans = append(ans, make([]int, len(next)))

		tmp := make([]*TreeNode, 0)

		for i, node := range next {
			if depth%2 == 1 {
				ans[depth][len(next)-i-1] = node.Val
			} else {
				ans[depth][i] = node.Val
			}
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
