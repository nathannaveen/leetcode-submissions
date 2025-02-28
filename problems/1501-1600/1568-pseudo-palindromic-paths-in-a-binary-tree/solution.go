/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func pseudoPalindromicPaths (root *TreeNode) int {
    res = 0
    dfs(root, 0, map[int] int{})
    return res
}

func dfs(root *TreeNode, cnt int, x map[int] int) {
    if root == nil {
        return
    }

    x[root.Val]++

    if x[root.Val] % 2 == 0 {
        cnt--
    } else {
        cnt++
    }

    if root.Left == nil && root.Right == nil {
        if cnt <= 1 {
            res++
        }
        return
    }

    copyX := map[int] int {}
    for k, v := range x {
        copyX[k] = v
    }

    dfs(root.Left, cnt, copyX)
    dfs(root.Right, cnt, x)
}