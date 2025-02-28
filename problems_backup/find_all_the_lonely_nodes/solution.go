/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getLonelyNodes(root *TreeNode) []int {
	stack := []*TreeNode{root}
	res := []int{}

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop.Left != nil && pop.Right == nil {
			res = append(res, pop.Left.Val)
			stack = append(stack, pop.Left)
		} else if pop.Left == nil && pop.Right != nil {
			res = append(res, pop.Right.Val)
			stack = append(stack, pop.Right)
		} else if pop.Left != nil && pop.Right != nil {
			stack = append(stack, pop.Left, pop.Right)
		}
	}
	return res
}