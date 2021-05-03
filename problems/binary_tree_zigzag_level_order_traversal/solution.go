func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	queue := []*TreeNode{root}
	fromRight := false

	for len(queue) != 0 {
		temp := []int{}
		n := len(queue)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				if fromRight {
					temp = append([]int{node.Val}, temp...)
				} else {
					temp = append(temp, node.Val)
				}
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}
		res = append(res, temp)
		fromRight = !fromRight
	}
	return res[:len(res)-1]
}
