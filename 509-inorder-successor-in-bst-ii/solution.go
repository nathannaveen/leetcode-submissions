/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func inorderSuccessor(node *Node) *Node {
	if node.Right != nil {
		node = node.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
	} else {
		val := node.Val
		node = node.Parent
		for node != nil && node.Val < val {
			node = node.Parent
		}
		return node
	}
}