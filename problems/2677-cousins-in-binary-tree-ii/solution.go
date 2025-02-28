/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func replaceValueInTree(root *TreeNode) *TreeNode {
    siblings := map[*TreeNode] int {} // node : sibling value
    
    queue := []*TreeNode{ root }
    
    for len(queue) != 0 {
        n := len(queue)
        sum := 0
        newSiblings := map[*TreeNode] int {}
        
        for i := 0; i < n; i++ {
            p := queue[i]
            
            sum += p.Val
            
            if p.Right != nil && p.Left != nil {
                newSiblings[p.Right] = p.Left.Val
                newSiblings[p.Left] = p.Right.Val
            }
        }
        
        for i := 0; i < n; i++ {
            pop := queue[0]
            queue = queue[1:]
            
            pop.Val = sum - pop.Val - siblings[pop]
            
            if pop.Right != nil {
                queue = append(queue, pop.Right)
            }
            
            if pop.Left != nil {
                queue = append(queue, pop.Left)
            }
        }
        
        siblings = newSiblings
    }
    
    return root
}