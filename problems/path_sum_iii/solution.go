/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type key struct {
    n *TreeNode
    arr []int
}

func pathSum(root *TreeNode, targetSum int) int {
    res := 0
    
    if root != nil {
        stack := []key{ key{root, []int{root.Val} } }
        if root.Val == targetSum {
            res++
        }

        for len(stack) != 0 {
            pop := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]

            if pop.n.Left != nil {
                arr := make([]int, len(pop.arr))
                copy(arr, pop.arr)
                arr = append(arr, 0)
                for i := range arr {
                    arr[i] += pop.n.Left.Val
                    if arr[i] == targetSum {
                        res++
                    }
                }
                stack = append(stack, key{pop.n.Left, arr})
            }

            if pop.n.Right != nil {
                arr := make([]int, len(pop.arr))
                copy(arr, pop.arr)
                arr = append(arr, 0)
                for i := range arr {
                    arr[i] += pop.n.Right.Val
                    if arr[i] == targetSum {
                        res++
                    }
                }
                stack = append(stack, key{pop.n.Right, arr})
            }
        }
    }
    
    return res
}