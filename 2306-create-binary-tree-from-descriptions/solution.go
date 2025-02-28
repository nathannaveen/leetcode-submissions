/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func createBinaryTree(descriptions [][]int) *TreeNode {
    child := map[int] bool {}
    left := map[int] int {} // map[parent] left child
    right := map[int] int {} // map[parent] right child
    
    for _, d := range descriptions {
        if d[2] == 1 { // left child
            left[d[0]] = d[1]
        } else { // right child
            right[d[0]] = d[1]
        }

        if !child[d[0]] {
            child[d[0]] = false
        }
        child[d[1]] = true
    }

    root := &TreeNode{}

    for a, b := range child {
        if !b {
            root = &TreeNode{Val: a}
        }
    }
    
    stack := []*TreeNode{ root }

    for len(stack) != 0 {
        pop := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]

        if _, ok := left[pop.Val]; ok {
            pop.Left = &TreeNode{Val : left[pop.Val]}
            stack = append(stack, pop.Left)
        }

        if _, ok := right[pop.Val]; ok {
            pop.Right = &TreeNode{Val : right[pop.Val]}
            stack = append(stack, pop.Right)
        }
    }

    return root
}