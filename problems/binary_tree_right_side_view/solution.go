/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    return helper([]int{}, []*TreeNode{root})
}

func helper(res []int, queue []*TreeNode ) []int {
    
    if queue[len(queue) - 1] != nil {
        n := len(queue)
        res = append(res, queue[len(queue) - 1].Val)

        for i := 0; i < n; i++ {
            pop := queue[0]
            queue = queue[1:]

            if pop.Left != nil {
                queue = append(queue, pop.Left)
            }
            if pop.Right != nil {
                queue = append(queue, pop.Right)
            }
        }

        if len(queue) != 0 {
            res = helper(res, queue)
        }
    }
    
    return res
}


// func rightSideView(root *TreeNode) []int {
// 	res := []int{}
//     if root != nil {
//         queue := []*TreeNode{ root }

//         for len(queue) != 0 {
//             n := len(queue)
//             res = append(res, queue[n - 1].Val)

//             for i := 0; i < n; i++ {
//                 pop := queue[0]
//                 queue = queue[1:]

//                 if pop.Left != nil {
//                     queue = append(queue, pop.Left)
//                 }
//                 if pop.Right != nil {
//                     queue = append(queue, pop.Right)
//                 }
//             }
//         }
//     }
    
// 	return res
// }