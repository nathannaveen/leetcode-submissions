/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseOddLevels(root *TreeNode) *TreeNode {
    queue := []*TreeNode{ root }
    c := 0
    
    for len(queue) != 0 {
        n := len(queue)
        arr := []int{}
        
        for i := 0; i < n; i++ {
            if queue[i] != nil {
                arr = append(arr, queue[i].Val)
            }
        }
        
        for i := 0; i < n; i++ {
            pop := queue[0]
            queue = queue[1:]
            
            if pop != nil {
                if c % 2 == 1 {
                    pop.Val = arr[len(arr) - i - 1]
                }
                
                queue = append(queue, pop.Left, pop.Right)
            }
        }
        
        c++
    }
    
    return root
}