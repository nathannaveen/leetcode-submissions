/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
    res := [][]int{}
    queue := []*Node{ root }
    
    if root == nil {
        return res
    }
    
    for len(queue) != 0 {
        n := len(queue)
        arr := []int{}
        
        for i := 0; i < n; i++ {
            pop := queue[0]
            queue = queue[1:]
            
            if pop == nil { continue }
            
            arr = append(arr, pop.Val)
            queue = append(queue, pop.Children...)
        }
        
        res = append(res, arr)
    }
    
    return res
}