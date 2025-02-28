/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Prev *Node
 * }
 */

func toArray(head *Node) []int {
    var arr []int

    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }

    return arr
}
