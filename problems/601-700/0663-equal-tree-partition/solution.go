/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res = false

func checkEqualTree(root *TreeNode) bool {
    res = false
    sum := traverse(root)

    traverse2(root.Left, sum)
    traverse2(root.Right, sum)

    return res
}

func traverse(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return root.Val + traverse(root.Left) + traverse(root.Right)
}

func traverse2(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }

    x := root.Val + traverse2(root.Left, sum) + traverse2(root.Right, sum)

    if sum - x == x {
        res = true
    }

    return x
}
