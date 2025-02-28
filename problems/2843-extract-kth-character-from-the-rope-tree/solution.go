/**
 * Definition for a rope tree node.
 * type RopeTreeNode struct {
 * 	   len   int
 * 	   val   string
 * 	   left  *RopeTreeNode
 * 	   right *RopeTreeNode
 * }
 */
func getKthCharacter(root *RopeTreeNode, k int) byte {
    return dfs(root)[k - 1]
}

func dfs(root *RopeTreeNode) string {
    if root == nil {
        return ""
    }
    if root.val != "" {
        return root.val
    }

    return dfs(root.left) + dfs(root.right)
}
