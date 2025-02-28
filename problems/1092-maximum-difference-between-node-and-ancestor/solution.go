/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func maxAncestorDiff(root *TreeNode) int {
    res = 0
    dfs(root)
    return res
}

func dfs(root *TreeNode) (int, int) {
    // return the max value and the min value

    rv := root.Val

    lMax, lMin := rv, rv
    rMax, rMin := rv, rv

    if root.Left != nil {
        lMax, lMin = dfs(root.Left)
    }
    if root.Right != nil {
        rMax, rMin = dfs(root.Right)
    }

    totalMax := max(max(lMax, rMax), rv)
    totalMin := min(min(lMin, rMin), rv)

    if abs(rv - totalMax) > res {
        res = abs(rv - totalMax)
    }
    if abs(rv - totalMin) > res {
        res = abs(rv - totalMin)
    }

    return totalMax, totalMin
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
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