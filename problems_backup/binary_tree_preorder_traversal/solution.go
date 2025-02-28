/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    if root != nil {
        return helper([]int{}, root)
    }
    return nil
}

func helper(res []int, root *TreeNode) []int {
    res = append(res, root.Val)
    if root.Left != nil {
        res = helper(res, root.Left)
    } 
    if root.Right != nil {
        res = helper(res, root.Right)
    }
    return res
}