func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		n := len(queue)
		var temp []int

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			
			if node != nil {
				temp = append(temp, node.Val)
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}
		res = append(res, temp)
	}
	return res[:len(res)-1]
}