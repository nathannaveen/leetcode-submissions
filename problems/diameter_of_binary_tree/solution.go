/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var maximum = 0

func diameterOfBinaryTree(root *TreeNode) int {
    maximum = 0
    helper(root)
    return maximum
}

func helper(root *TreeNode) int {
    if root == nil { return -1 }
    
    leftHeight, rightHeight := helper(root.Left), helper(root.Right)
    
    maximum = max(leftHeight + rightHeight + 2, maximum)
    
    return 1 + max(leftHeight, rightHeight)
}

func max(a, b int) int {
    if a > b { return a }
    return b
}