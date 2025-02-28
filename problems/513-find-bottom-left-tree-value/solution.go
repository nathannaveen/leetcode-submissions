/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findBottomLeftValue(root *TreeNode) int {
	queue := []*TreeNode{root}
	bottomLeft := 0

	for len(queue) != 0 {
		n := len(queue)
		bottomLeft = queue[0].Val
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	
	return bottomLeft
}