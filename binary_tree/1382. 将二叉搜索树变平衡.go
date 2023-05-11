package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func balanceBST(root *TreeNode) *TreeNode {
	var arr = []int{}

	var f func(*TreeNode)

	// 这里中序遍历
	// 中序遍历之后，获得有序的数组，这时候每次取mid=(left+right)/2来做子树的根节点
	// 一定是平衡的
	f = func(n *TreeNode) {
		if n == nil {
			return
		}
		f(n.Left)
		arr = append(arr, n.Val)
		f(n.Right)
	}

	f(root)

	// fmt.Println(arr)

	var f1 func([]int, int, int) *TreeNode
	f1 = func(nums []int, left, right int) *TreeNode {

		if left > right {
			return nil
		}

		mid := (left + right) / 2
		node := &TreeNode{
			Val: nums[mid],
		}

		node.Left = f1(nums, left, mid-1)
		node.Right = f1(nums, mid+1, right)

		return node

	}

	return f1(arr, 0, len(arr)-1)
}
