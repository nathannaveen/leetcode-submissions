/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfSubtree(root *TreeNode) int {
    _, _, res := dfs(root)
    return res
}

func dfs(root *TreeNode) (int, int, int) {
    /*
    This function returns:
    
    * the sum of all elements in the subtree 
    * the number of nodes in the subtree
    * and the result for the subtree
    */
    
    if root == nil {
        return 0, 0, 0
    }
    
    sum1, n1, res1 := dfs(root.Left)
    sum2, n2, res2 := dfs(root.Right)
    
    sum := sum1 + sum2 + root.Val
    n := n1 + n2 + 1
    res := res1 + res2
    
    if sum / n == root.Val {
        res++
    }
    
    return sum, n, res
}