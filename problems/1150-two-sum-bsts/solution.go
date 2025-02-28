/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	stack := []*TreeNode{root1}
	m := make(map[int]int)

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop != nil {
			stack = append(stack, pop.Left, pop.Right)
			m[target-pop.Val] = pop.Val
		}
	}

	stack = append(stack, root2)

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop != nil {
			stack = append(stack, pop.Left, pop.Right)
			if _, ok := m[pop.Val]; ok {
				return true
			}
		}
	}

	return false
}