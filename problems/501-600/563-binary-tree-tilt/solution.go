/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res = 0

func findTilt(root *TreeNode) int {
    res = 0
    helper(root)
    return res
}

func helper(node *TreeNode) int {
    if node == nil { return 0 }
    
    left := helper(node.Left)
    right := helper(node.Right)
    
    res += int(math.Abs(float64(left - right)))
    
    return left + right + node.Val
}