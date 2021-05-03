func isEvenOddTree(root *TreeNode) bool {
	row := 0
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		n := len(queue)
		for i2 := 0; i2 < n; i2++ {
			node := queue[0]
			queue = queue[1:]
			if row % 2 == 0 && (node.Val % 2 == 0 || (i2 != n - 1 && node.Val >= queue[0].Val)) {
				return false
			}
			if row % 2 == 1 && (node.Val % 2 == 1 || (i2 != n - 1 && node.Val <= queue[0].Val)) {
				return false
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		row++
	}
	return true
}