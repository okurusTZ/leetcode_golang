package binary_tree

func maxAncestorDiff(root *TreeNode) int {
	ans := 0

	var f func(*TreeNode, int, int)

	f = func(node *TreeNode, minimum int, maximum int) {
		if node == nil {
			return
		}

		// 记录下递归的时候最大值和最小值
		minimum = min(minimum, node.Val)
		maximum = max(maximum, node.Val)

		ans = max(ans, maximum-minimum)

		f(node.Left, minimum, maximum)
		f(node.Right, minimum, maximum)
	}

	f(root, 10^5+1, -1)

	return ans

}
