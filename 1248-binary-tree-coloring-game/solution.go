/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        pop := queue[0]
        queue = queue[1:]
        
        if pop == nil {
            continue
        }

        if pop.Val == x {
            left := dfs(pop.Left)
            right := dfs(pop.Right)
            rest := n - (left + right + 1)

            if rest > left + right + 1 || 
                left > right + 1 + rest || 
                left + 1 + rest < right {
                return true
            } else {
                return false
            }
        }

        queue = append(queue, pop.Left, pop.Right)
    }

    return false
}

func dfs(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + dfs(root.Left) + dfs(root.Right)
}