/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    stack := []*TreeNode{root}
    
    for len(stack) != 0 {
        pop := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        
        if pop != nil && subRoot != nil {
            if pop.Val == subRoot.Val && isSame(pop, subRoot) {
                return true
            }
            stack = append(stack, pop.Right, pop.Left)
        }
    }
    
    return false
}

func isSame(root *TreeNode, root2 *TreeNode) bool {
    if root != nil && root2 != nil {
        if root.Val == root2.Val && isSame(root.Right, root2.Right) && isSame(root.Left, root2.Left) {
            return true
        }
        return false
    } else if root == nil && root2 == nil {
        return true
    } else {
        return false
    }
}