/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    previous := -2147483649
    stack := []*TreeNode{}
    
    for len(stack) != 0 || root != nil {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        
        if stack[len(stack) - 1].Val <= previous { 
            return false
        }
        previous = stack[len(stack) - 1].Val
        root = stack[len(stack) - 1].Right
        stack = stack[:len(stack) - 1]
    }
    
    return true
}