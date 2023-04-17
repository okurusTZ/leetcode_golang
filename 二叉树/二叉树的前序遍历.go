package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	// res := make([]int, 0)
	// if root == nil {
	// 	return res
	// }
	// res = preOrderRecur(root)
	// return res

	stack := []*TreeNode{} // 模拟一个栈
	node := root
	val := make([]int, 0)

	for len(stack) > 0 || node != nil {
		for node != nil {
			val = append(val, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return val
}

func preOrderRecur(head *TreeNode, res []int) []int {
	if head == nil {
		return res
	}
	res = append(res, head.Val)
	res = preOrderRecur(head.Left, res)
	res = preOrderRecur(head.Right, res)
	return res
}

func main() {}
