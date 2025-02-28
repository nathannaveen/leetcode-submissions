func isCompleteTree(root *TreeNode) bool {
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		isNextRow := false
		n := len(queue)

		if queue[0] != nil && queue[0].Left != nil {
			isNextRow = true
		}

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if isNextRow && node == nil {
				return false
			}
			if node == nil && len(queue) != 0 && queue[0] != nil {
				return false
			}
			if node != nil {
				queue = append(queue, node.Left, node.Right)
			}
		}
	}
	return true
}