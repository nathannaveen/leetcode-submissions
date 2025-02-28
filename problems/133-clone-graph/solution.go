/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    if node == nil {
        return nil
    }
    return helper(node, make(map[*Node] *Node))
}

func helper(node *Node, m map[*Node] *Node) *Node {
    if val, ok := m[node]; ok {
        return val
    }
    
    newNode := &Node{Val: node.Val}
    m[node] = newNode
    
    for _, neighbor := range node.Neighbors {
        newNode.Neighbors = append(newNode.Neighbors, helper(neighbor, m))
    }
    
    return newNode
}