/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	arr := []int{}
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) != 0 {
		g := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if g != nil {
			arr = append(arr, g.Val)
			stack = append(stack, g.Left, g.Right)
		}
	}

	sort.Ints(arr)

	return arr[k - 1]
}