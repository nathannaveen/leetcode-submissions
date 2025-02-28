/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    return helper(root, []int{})
}

func helper(r *Node, l []int)[]int{
    if r != nil{
        for _,item := range r.Children{
            l = helper(item,l)
        }
        l = append(l,r.Val)        
    }
    return l
}