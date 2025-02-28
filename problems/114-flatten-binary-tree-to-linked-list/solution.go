func flatten(root *TreeNode)  {
    if root != nil {
        stack := []*TreeNode{root}
        arr := []int{}

        for len(stack) != 0 {
            node := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]

            if node != nil {
                arr = append(arr, node.Val)
                stack = append(stack, node.Right, node.Left)
            }
        }

        tree := root

        tree.Left = nil
        tree.Right = nil 

        for i := 1; i < len(arr); i++ {
            tree.Right = &TreeNode{Val: arr[i],}
            tree = tree.Right
        }
    }
}