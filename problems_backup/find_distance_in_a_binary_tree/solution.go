/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var P, Q = 0, 0
func findDistance(root *TreeNode, p int, q int) int {
    P = p
    Q = q
    
    ancestor := dfs(root)
    return dist(ancestor, p) + dist(ancestor, q)
}

func dist(root *TreeNode, t int) int {
    if root == nil {
        return 10000
    }
    if root.Val == t {
        return 0
    }
    return 1 + int(math.Min(float64(dist(root.Left, t)), float64(dist(root.Right, t))))
}

func dfs(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    left := dfs(root.Left)
    right := dfs(root.Right)
    
    if root.Val == P || root.Val == Q || (left != nil && right != nil) {
        return root
    }
    
    if left == nil {
        return right
    }
    return left
}