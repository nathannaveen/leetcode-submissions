/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	stack := []*TreeNode{root1}
	arr1 := []int{}
	arr2 := []int{}

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop != nil {
			stack = append(stack, pop.Left, pop.Right)
			arr1 = append(arr1, pop.Val)
		}
	}

	stack = append(stack, root2)

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop != nil {
			stack = append(stack, pop.Left, pop.Right)
			arr2 = append(arr2, pop.Val)
		}
	}

	for _, i2 := range arr1 {
		for _, i4 := range arr2 {
			if i2+i4 == target {
				return true
			}
		}
	}
	return false
}
