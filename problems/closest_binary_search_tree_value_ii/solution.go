/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func closestKValues(root *TreeNode, target float64, k int) []int {
	h := []int{}
	g := []float64{}
	stack := []*TreeNode{root}

	for len(stack) != 0 {
		pop := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		if pop != nil {
			g = append(g, abs(float64(pop.Val) - target))
			h = append(h, pop.Val)
			stack = append(stack, pop.Left, pop.Right)
		}
	}

	for i := 1; i < len(g); i++ {
		if i >= 1 && g[i-1] > g[i] {
			h[i], h[i - 1] = h[i - 1], h[i]
			g[i], g[i - 1] = g[i - 1], g[i]
			i -= 2
		}
	}
	
	return h[:k]
}

func abs(a float64) float64 {
	if a > 0 {
		return a
	}
	return -a
}