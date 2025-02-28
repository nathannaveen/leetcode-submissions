/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil { return 0 }
    left, right := maxDepth(root.Left), maxDepth(root.Right)
    
    return 1 + int(math.Max(float64(left), float64(right)))
}