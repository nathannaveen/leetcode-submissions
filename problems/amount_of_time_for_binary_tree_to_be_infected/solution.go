/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
    connections := map[int] []int {}
    dfs(connections, root)
    
    visited := map[int] bool {}
    queue := []int{ start }
    
    c := 0
    
    for len(queue) != 0 {
        n := len(queue)
        
        done := false
        
        for i := 0; i < n; i++ {
            pop := queue[0]
            queue = queue[1:]
            
            if !visited[pop] {
                queue = append(queue, connections[pop]...)
                done = true
            }
            visited[pop] = true
        }
        
        if done {
            c++
        }
    }
    
    return c - 1
}

func dfs(connections map[int] []int, node *TreeNode) {
    if node == nil {
        return
    }
    
    if node.Left != nil {
        connections[node.Val] = append(connections[node.Val], node.Left.Val)
        connections[node.Left.Val] = append(connections[node.Left.Val], node.Val)
        dfs(connections, node.Left)
    }
    
    if node.Right != nil {
        connections[node.Val] = append(connections[node.Val], node.Right.Val)
        connections[node.Right.Val] = append(connections[node.Right.Val], node.Val)
        dfs(connections, node.Right)
    }
}