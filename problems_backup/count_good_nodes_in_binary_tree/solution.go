/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type key struct {
    n *TreeNode
    max int
}

func goodNodes(root *TreeNode) int {
    stack := []key{ key{ root, root.Val } }
    res := 0
    
    for len(stack) != 0 {
        node := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        
        if node.n != nil {
            if node.n.Val >= node.max {
                res++
                node.max = node.n.Val
            }
            
            stack = append(stack, 
                           key{node.n.Right, node.max}, 
                           key{node.n.Left, node.max})
        }
    }
    
    return res
}