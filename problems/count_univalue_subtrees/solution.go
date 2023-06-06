/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res int

func countUnivalSubtrees(root *TreeNode) int {
    res = 0
    dfs(root)
    return res
}

func dfs(root *TreeNode) map[int] int {
    if root == nil {
        return map[int] int {}
    }

    a := dfs(root.Left)
    b := dfs(root.Right)

    for k, v := range b {
        a[k] = v
    }
    a[root.Val]++

    if len(a) == 1{
        res++
    }

    return a
}