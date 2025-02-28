func maxLevelSum(root *TreeNode) int {
	maximum, minLevel, level := root.Val, 1, 0
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		sum, n := 0, len(queue)

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		level++
		if sum > maximum {
			maximum, minLevel = sum, level
		}
	}
	return minLevel
}