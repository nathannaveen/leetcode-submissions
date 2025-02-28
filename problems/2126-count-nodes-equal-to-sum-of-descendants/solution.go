/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int = 0

func equalToDescendants(root *TreeNode) int {
    res = 0
    helper(root)
    return res
}

func helper(root *TreeNode) int {
    sum := 0
    if root.Right != nil {
        sum += helper(root.Right)
    }

    if root.Left != nil {
        sum += helper(root.Left)
    }

    if root.Val == sum {
        res++
    }

    return root.Val + sum
}