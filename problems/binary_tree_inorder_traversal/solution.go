/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    stack := []*TreeNode{}
    res := []int{}
    
    for root != nil || len(stack) != 0 {
        if root != nil {
            stack = append(stack, root)
            root = root.Left
        } else {
            root = stack[len(stack) - 1]
            root = root.Right
            res = append(res, stack[len(stack) - 1].Val)
            stack = stack[:len(stack) - 1]
        }
    }
    
    return res
}