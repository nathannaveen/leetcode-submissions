func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) != 0 {
		n := len(queue)
		temp := []int{}
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				temp = append(temp, node.Val)
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
		}
		if len(temp) != 0 {
			res = append(res, temp)
		}
	}

	left, right := 0, len(res)-1

	for left < right {
		res[left], res[right] = res[right], res[left]
		left++
		right--
	}
	
	return res
}