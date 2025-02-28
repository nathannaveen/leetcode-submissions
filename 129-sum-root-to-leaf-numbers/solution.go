/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type key struct {
    node *TreeNode
    val int
}

func sumNumbers(root *TreeNode) int {
    res := 0
    stack := []key{}
    stack = append(stack, key{ root, root.Val })
    
    for len(stack) != 0 {
        pop := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        
        if pop.node != nil {
            if pop.node.Left == nil && pop.node.Right == nil {
                res += pop.val
            } else {
                if pop.node.Left != nil {
                    stack = append(stack, key{pop.node.Left, 
					(pop.val) * 10 + pop.node.Left.Val })
                }
                if pop.node.Right != nil {
                    stack = append(stack, key{pop.node.Right, 
					(pop.val) * 10 + pop.node.Right.Val })
                }
            }
        }
    }
    
    return res
}