func isCousins(root *TreeNode, x int, y int) bool {
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		n := len(queue)
		containsX, containsY := false, false

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Right != nil && node.Left != nil && 
				((node.Right.Val == x && node.Left.Val == y) ||
				(node.Right.Val == y && node.Left.Val == x)) {

				return false
			}
			if node.Val == x {
				containsX = true
			} else if node.Val == y {
				containsY = true
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
		}

		if containsX && containsY {
			return true
		} else if (containsX && !containsY) || (containsY && !containsX) {
			return false
		}
	}

	return false
}