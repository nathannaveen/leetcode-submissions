/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    return helper(strconv.Itoa(root.Val), root, []string{})
}

func helper(s string, r *TreeNode, res []string) []string {
    
    if r.Left == nil && r.Right == nil {
        res = append(res, s)
    }
    if r.Left != nil {
        res = helper(s + "->" + strconv.Itoa(r.Left.Val), r.Left, res)
    }
    if r.Right != nil {
        res = helper(s + "->" + strconv.Itoa(r.Right.Val), r.Right, res)
    }
    
    return res
}