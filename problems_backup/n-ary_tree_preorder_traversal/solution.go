/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    return helper(root, []int{})
}

func helper(r *Node, l []int)[]int{
    if r != nil{
        l = append(l,r.Val)
        for _,item := range r.Children{
            l = helper(item,l)
        }
    }
    return l
}