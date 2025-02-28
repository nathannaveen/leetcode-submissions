/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var minimum *TreeNode
	stack := []*TreeNode{root}

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop != nil {
			if minimum == nil && pop.Val > p.Val {
				minimum = pop
			} else if minimum != nil && pop.Val > p.Val && minimum.Val > pop.Val {
				minimum = pop
			}

			stack = append(stack, pop.Left, pop.Right)
		}
	}
	return minimum
}