/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	deapest := 0
	h := []*TreeNode{root}

	for len(h) > 0 {
		sum := 0
		n := len(h)

		for i := 0; i < n; i++ {
			node := h[0]
			h = h[1:]
			sum += node.Val

			if node.Left != nil {
				h = append(h, node.Left)
			}
			if node.Right != nil {
				h = append(h, node.Right)
			}
		}
		deapest = sum
	}
	return deapest
}