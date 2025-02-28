/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res float64 = float64(0)

type key struct {
    counter int
    val int
}

func maximumAverageSubtree(root *TreeNode) float64 {
    res = float64(0)
    helper(root)
    return res
}

func helper(root *TreeNode) key {
    sum := root.Val
    counter := 1
    
    if root.Right != nil {
        temp := helper(root.Right)
        sum += temp.val
        counter += temp.counter
    }
    
    if root.Left != nil {
        temp := helper(root.Left)
        sum += temp.val
        counter += temp.counter
    }
    
    res = math.Max(res, float64(sum) / float64(counter))
    
    return key{ counter, sum }
}