/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	h := []*TreeNode{root}
	res := []float64{}

	for len(h) != 0 {
		n, sum := len(h), float64(0)

		for i := 0; i < n; i++ {
			cur := h[0]
			h = h[1:]
			sum += float64(cur.Val)
			if cur.Left != nil {
				h = append(h, cur.Left)
			}
			if cur.Right != nil {
				h = append(h, cur.Right)
			}
		}
		
		res = append(res, sum / float64(n))
	}

	return res
}