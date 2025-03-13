/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res = []int{}

func kthLargestPerfectSubtree(root *TreeNode, k int) int {
    res = []int{}
    helper(root)
    if k > len(res) {
        return -1
    }
    sort.Slice(res, func(i, j int) bool {
        return res[i] > res[j]
    })
    return res[k - 1]
}

func helper(root *TreeNode) (bool, int, int) {
    if root == nil {
        return true, 0, 0
    }

    a, b, e := helper(root.Left)
    c, d, f := helper(root.Right)

    if !a || !c {
        return false, 0, 0
    }
    if b != d {
        return false, 0, 0
    }

    res = append(res, e + f + 1)

    return true, b + 1, e + f + 1
}
