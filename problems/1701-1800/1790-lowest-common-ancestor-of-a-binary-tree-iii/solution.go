/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

func lowestCommonAncestor(p *Node, q *Node) *Node {
    m := make(map[*Node] bool)
    for p != nil {
        m[p] = true
        p = p.Parent
    }
    for q != nil {
        if _, ok := m[q]; ok {
            return q
        }
        q = q.Parent
    }
    return nil
}