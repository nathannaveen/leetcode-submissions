/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res = -1001

func maxPathSum(root *TreeNode) int {
    res = -1001
    helper(root)
    return res
}

func helper(root *TreeNode) int {
    if root == nil { return 0 }
    
    one := helper(root.Left)
    two := helper(root.Right)
    
    if one < 0 { one = 0 }
    if two < 0 { two = 0 }
    
    if res < root.Val + one + two {
        res = root.Val + one + two
    }
    
    return int(math.Max(float64(one), float64(two))) + root.Val
}
