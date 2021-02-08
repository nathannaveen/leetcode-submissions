/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	stack := []*TreeNode{}
	stack = append(stack, root)
	counter := 0
	for len(stack) > 0 {
		node := &TreeNode{}
		node, stack = stack[len(stack) - 1], stack[:len(stack) - 1]
		if node != nil {
			counter++
			stack = append(stack, node.Left, node.Right)
		}
	}
	return counter
}