package binary_tree

// 找到最近的公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}

	return nil

}

// 找二叉搜索树的公共祖先
// 利用二叉搜索树性质 left < root < right

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil || root == p || root == q {
		return root
	}

	// 这里应该比较当前节点值
	val := root.Val
	if min(p.Val, q.Val) > val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if max(p.Val, q.Val) < val {
		return lowestCommonAncestor(root.Left, p, q)
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}
	return nil
}
