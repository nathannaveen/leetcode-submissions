/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	queue := []*TreeNode{root}
	res := []int{}

	if root == nil {
		return res
	}

	for len(queue) != 0 {
		n := len(queue)
		maximum := queue[0].Val
		
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Val > maximum {
				maximum = node.Val
			}
			if node.Left != nil { queue = append(queue, node.Left) }
			if node.Right != nil { queue = append(queue, node.Right) }
		}
		res = append(res, maximum)
	}

	return res
}