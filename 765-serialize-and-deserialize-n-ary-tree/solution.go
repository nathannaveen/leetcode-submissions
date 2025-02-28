/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Codec struct {
    node *Node
}

func Constructor() *Codec {
    return &Codec{}
}

func (this *Codec) serialize(root *Node) string {
    this.node = root
    return "taco"
}

type key struct {
    one *Node
    two *Node
}

func (this *Codec) deserialize(data string) *Node {
    if this.node != nil {
        res := &Node{ Val: this.node.Val }
        stack := []key{ key{ this.node, res } }

        for len(stack) != 0 {
            pop := stack[len(stack) - 1]
            stack = stack[: len(stack) - 1]

            for _, i := range pop.one.Children {
                pop.two.Children = append(pop.two.Children, &Node{ Val: i.Val })
                stack = append(stack, key{ i, pop.two.Children[len(pop.two.Children) - 1] } )
            }
        }
        return res
    }
    return nil
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */