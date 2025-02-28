/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res = true

func isBalanced(root *TreeNode) bool {
    res = true
    dfs(root)
    return res
}

func dfs(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left, right := dfs(root.Left), dfs(root.Right)
    
    if abs(left - right) > 1 {
        res = false
    }

    return 1 + max(left, right)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}