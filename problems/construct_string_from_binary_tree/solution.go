func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	var str []string
	preorderStr(t, &str)

	s := strings.Join(str, "")
	s = strings.TrimPrefix(s, "(")
	s = strings.TrimSuffix(s, ")")
	return s
}

func preorderStr(t *TreeNode, str *[]string) {
	*str = append(*str, "(", strconv.Itoa(t.Val))

	if t.Left == nil && t.Right == nil {
		*str = append(*str, ")")
		return
	}

	if t.Left == nil {
		*str = append(*str, "()")
	} else {
		preorderStr(t.Left, str)
	}

	if t.Right != nil {
		preorderStr(t.Right, str)
	}
	*str = append(*str, ")")
}