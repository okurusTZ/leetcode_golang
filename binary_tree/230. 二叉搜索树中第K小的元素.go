package binary_tree

func kthSmallest(root *TreeNode, k int) int {
	// 中序遍历找到第k个

	var c = 0
	var res = 0

	var f func(*TreeNode)
	f = func(tn *TreeNode) {
		if tn == nil {
			return
		}

		f(tn.Left)
		c++
		if c == k {
			res = tn.Val
			return
		}
		f(tn.Right)
	}

	f(root)
	return res
}
