/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	stack = append(stack, root1)

	for len(stack) != 0 {
		g := &TreeNode{}
		g, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if g != nil {
			res = append(res, g.Val)
			stack = append(stack, g.Right, g.Left)
		}
	}

	stack = append(stack, root2)

	for len(stack) != 0 {
		g := &TreeNode{}
		g, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if g != nil {
			res = append(res, g.Val)
			stack = append(stack, g.Right, g.Left)
		}
	}	
	
	sort.Ints(res)
	return res
}