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
    col int
}

func verticalOrder(root *TreeNode) [][]int {
    m := map[int] []int {} // cols, nodes per column
    queue := []key{ key{ root, 0 } }
    min := 0 // minimum column
    
    for len(queue) != 0 {
        n := len(queue)
        
        for i := 0; i < n; i++ {
            pop := queue[0]
            queue = queue[1:]
            
            if pop.node == nil {
                continue
            }
            if pop.col < min {
                min = pop.col
            }
            
            m[pop.col] = append(m[pop.col], pop.node.Val)
            
            queue = append(queue, 
                key{ pop.node.Left, pop.col - 1 }, 
                key{ pop.node.Right, pop.col + 1 },
            )
        }
    }
    
    min *= -1
    
    res := make([][]int, len(m))
    
    for a, b := range m {
        res[min + a] = b
    }
    
    return res
}