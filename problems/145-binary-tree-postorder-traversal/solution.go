/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) != 0 {
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if val != nil {
			res = append([]int{val.Val}, res...)
			stack = append(stack, val.Left, val.Right)
		}
	}
	return res
}