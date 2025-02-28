/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return root
	}
	if shouldDelete(root, target){
		return nil
	}
	return root
}

func shouldDelete(root *TreeNode, target int) bool {

	if root.Left != nil && shouldDelete(root.Left, target) {
		root.Left = nil
	}

	if root.Right != nil && shouldDelete(root.Right, target) {
		root.Right = nil
	}
	return root.Right == nil && root.Left == nil && root.Val == target
}