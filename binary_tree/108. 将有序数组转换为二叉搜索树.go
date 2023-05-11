package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	// 可以看作是中序排列

	var f func([]int, int, int) *TreeNode
	f = func(nums []int, left, right int) *TreeNode {

		if left > right {
			return nil
		}

		mid := (left + right) / 2
		node := &TreeNode{
			Val: nums[mid],
		}

		node.Left = f(nums, left, mid-1)
		node.Right = f(nums, mid+1, right)

		return node

	}

	return f(nums, 0, len(nums)-1)
}
