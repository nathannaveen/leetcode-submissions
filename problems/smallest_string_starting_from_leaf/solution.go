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
    val string
}

func smallestFromLeaf(root *TreeNode) string {
    res := []string{}
    stack := []key{ key{ root, string(root.Val + 'a') } } 
    
    for len(stack) != 0 {
        pop := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        
        if pop.node != nil {
            if pop.node.Left == nil && pop.node.Right == nil {
                res = append(res, pop.val)
            } else {
                if pop.node.Left != nil {
                    stack = append(stack, key{pop.node.Left, 
                    string(pop.node.Left.Val + 'a') + pop.val })
                }
                if pop.node.Right != nil {
                    stack = append(stack, key{pop.node.Right, 
					string(pop.node.Right.Val + 'a') + pop.val })
                }
            }
        }
    }
    
    sort.Strings(res)
    
    return res[0]
}