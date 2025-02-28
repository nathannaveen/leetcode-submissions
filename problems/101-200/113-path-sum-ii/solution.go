/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type key struct {
    node *TreeNode
    val int
    arr []int
}

func pathSum(root *TreeNode, targetSum int) [][]int {
    res := [][]int{}
    if root != nil {
        stack := []key{}
        stack = append(stack, key{ root, root.Val, []int{root.Val} })

        for len(stack) != 0 {
            pop := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]

            if pop.node != nil {
                if pop.node.Left == nil && pop.node.Right == nil && pop.val == targetSum {
                    res = append(res, pop.arr)
                } else {
                    if pop.node.Left != nil {
                        temp := []int{}
                        temp = append(temp, pop.arr...)
                        temp = append(temp, pop.node.Left.Val)
                        stack = append(stack, key{pop.node.Left, 
                        (pop.val) + pop.node.Left.Val, temp })
                    }
                    
                    if pop.node.Right != nil {
                        temp := []int{}
                        temp = append(temp, pop.arr...)
                        temp = append(temp, pop.node.Right.Val)
                        stack = append(stack, key{pop.node.Right, 
                        (pop.val) + pop.node.Right.Val, temp })
                    }
                }
            }
        }
    }
    return res
}