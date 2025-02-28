/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    return helper(nums)
}

func helper(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    
    index := 0
    
    for i, n := range nums {
        if n > nums[index] {
            index = i
        }
    }
    
    return &TreeNode{ nums[index], helper(nums[:index]), helper(nums[index + 1:]) }
}