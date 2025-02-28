/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
	max := 0
	m := make(map[int]int)
	res := []int{}
	stack := []*TreeNode{}
	stack = append(stack, root)

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pop != nil {
			m[pop.Val]++
			stack = append(stack, pop.Left, pop.Right)
		}
	}

	for i, i2 := range m {
		if i2 > max {
			res = []int{i}
			max = i2
		} else if i2 == max {
			res = append(res, i)
		}
	}
	return res
}