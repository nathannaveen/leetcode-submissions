/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
    queue := []*TreeNode{root}
    
    for len(queue) != 0 {
        n := len(queue)
        
        for i := 0; i < n; i++ {
            pop := queue[0] // first element
            queue = queue[1:] // actualy removing the first element
            
            if pop.Val == u.Val {
                if i == n - 1 {
                    return nil
                }
                return queue[0]
            }
            
            if pop.Left != nil {
                queue = append(queue, pop.Left)
            }
            if pop.Right != nil {
                queue = append(queue, pop.Right)
            }
        }
    }
    
    return nil
}