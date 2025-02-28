/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
    queue := []*TreeNode{ root }
    counter := 1
    
    if counter == depth {
        return &TreeNode{ Val : val, Left : root }
    }
    
    for len(queue) != 0 {
        n := len(queue)
        counter++
        
        if counter == depth {
            for i := 0; i < n; i++ {
                pop := queue[0]
                queue = queue[1 :]
                
                r, l := &TreeNode{ Val : val, Right : pop.Right }, &TreeNode{ Val : val, Left : pop.Left }
                pop.Right = r
                pop.Left = l
            }
            break
        } else {
            for i := 0; i < n; i++ {
                pop := queue[0]
                queue = queue[1 :]
                if pop.Left != nil {
                    queue = append(queue, pop.Left)
                }
                if pop.Right != nil {
                    queue = append(queue, pop.Right)
                }
            }
        }
    }
    return root
}