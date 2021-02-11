/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) != 0 {
		val := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if val != nil {
			res = append(res, val.Val)
			stack = append(stack, val.Right, val.Left)
		}
	}
	return res
}