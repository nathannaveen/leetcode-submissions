/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	queue := []*Node{ root }

    appendToQueue := func(node *Node) {
        if node != nil && node.Left != nil {
            queue = append(queue, node.Left, node.Right)
        }
    }

    for len(queue) > 0 {
        prev := queue[0]
        queue = queue[1:]

        n := len(queue)

        appendToQueue(prev)

        for i := 0; i < n; i++ {
            pop := queue[0]
            queue = queue[1:]

            prev.Next = pop

            appendToQueue(pop)

            prev = prev.Next
        }
    }

    return root
}