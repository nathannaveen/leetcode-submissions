/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	h := []*TreeNode{}
	if root == nil {
		root = &TreeNode{
			Val:   val,
		}
		return root
	}
	h = append(h, root)

	for len(h) != 0 {
		pop := &TreeNode{}
		pop, h = h[len(h)-1], h[:len(h)-1]

		if pop != nil {
			if val < pop.Val {
				if pop.Left == nil {
					pop.Left = &TreeNode{
						Val: val,
					}
					return root
				}
				h = append(h, pop.Left)
			} else {
				if pop.Right == nil {
					pop.Right = &TreeNode{
						Val: val,
					}
					return root
				}
				h = append(h, pop.Right)
			}
		}
	}

	return root
}