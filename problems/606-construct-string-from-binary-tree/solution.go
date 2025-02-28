func tree2str(t *TreeNode) string {
	stack := []*TreeNode{t}
	m := make(map[*TreeNode]int)
	res := ""

	for len(stack) != 0 {
		popedVal := stack[len(stack)-1]

		if popedVal != nil {
			if m[popedVal] == 1 {
				stack = stack[:len(stack)-1]
				res += ")"
			} else {
				m[popedVal] = 1
				res += "(" + strconv.Itoa(popedVal.Val)
				if popedVal.Left == nil && popedVal.Right != nil {
					res += "()"
				}
				stack = append(stack, popedVal.Right, popedVal.Left)
			}
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return res[1 : len(res)-1]
}