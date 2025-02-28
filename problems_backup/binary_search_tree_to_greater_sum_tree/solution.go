/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func bstToGst(root *TreeNode) *TreeNode {
	arr := []int{}
	total := 0
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop != nil {
			arr = append(arr, pop.Val)
			total += pop.Val
			stack = append(stack, pop.Left, pop.Right)
		}
	}

	sort.Ints(arr)

	stack = append(stack, root)

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop != nil {
			sum := 0
			for i := 0; i < len(arr); i++ {
				if arr[i] == pop.Val {
					break
				}
				sum += arr[i]
			}
			pop.Val = total - sum
			stack = append(stack, pop.Left, pop.Right)
		}
	}

	return root
}