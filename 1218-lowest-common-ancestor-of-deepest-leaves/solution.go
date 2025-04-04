/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
    _, res := dfs(root)

    return res
}

func dfs(root *TreeNode) (int, *TreeNode) {
    left := 0
    lNode := &TreeNode{}
    right := 0
    rNode := &TreeNode{}
    
    if root.Left != nil {
        left, lNode = dfs(root.Left)
    }
    if root.Right != nil {
        right, rNode = dfs(root.Right)
    }

    if left == right {
        return left + 1, root
    }
    if left > right {
        return left + 1, lNode
    }
    return right + 1, rNode
}
