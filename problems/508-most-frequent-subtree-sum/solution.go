/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res []int = []int{}
var m map[int] int = make(map[int] int)
var max int = 0

func findFrequentTreeSum(root *TreeNode) []int {
    res, m, max = []int{}, make(map[int] int), 0
    helper(root)
    return res
}

func helper(root *TreeNode) int {
    sum := root.Val
    
    if root.Right != nil {
        sum += helper(root.Right)
    }
    
    if root.Left != nil {
        sum += helper(root.Left)
    }
    
    m[sum]++
    
    if m[sum] > max {
        res = []int{ sum }
        max = m[sum]
    } else if m[sum] == max {
        res = append(res, sum)
    }
    
    return sum
}