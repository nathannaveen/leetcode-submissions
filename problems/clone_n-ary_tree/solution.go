/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type key struct {
    one *Node
    two *Node
}

func cloneTree(root *Node) *Node {
    if root != nil {
        res := &Node{ Val: root.Val, }
        stack := []key{ key{root, res} }

        for len(stack) != 0 {
            pop := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            
            for i := 0; i < len(pop.one.Children) && pop.one != nil; i++ {
                pop.two.Children = append(pop.two.Children, &Node{ Val: pop.one.Children[i].Val, })
                stack = append(stack, key{ pop.one.Children[i], pop.two.Children[i] })
            }
        }
        return res
    }
    return nil
}
