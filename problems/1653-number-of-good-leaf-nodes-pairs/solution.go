/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int

func countPairs(root *TreeNode, distance int) int {
    res = 0
    dfs(root, distance)
    return res
}

func dfs(root *TreeNode, dist int) []int {
    // return the distances 

    if root.Left == nil && root.Right == nil {
        return []int{0}
    }

    lefts, rights := []int{}, []int{}

    if root.Left != nil {
        lefts = dfs(root.Left, dist)

        for i := 0; i < len(lefts); i++ {
            lefts[i]++
        }
    }

    if root.Right != nil {
        rights = dfs(root.Right, dist)

        for i := 0; i < len(rights); i++ {
            rights[i]++
        }
    }

    for i := 0; i < len(lefts); i++ {
        for j := 0; j < len(rights); j++ {
            if lefts[i] + rights[j] <= dist {
                res++
            }
        }
    }

    return append(lefts, rights...)
}