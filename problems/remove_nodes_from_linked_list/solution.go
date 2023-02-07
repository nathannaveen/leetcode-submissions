/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes(head *ListNode) *ListNode {
    z := &ListNode{100001, head}
    cur := head
    prev := []*ListNode{z, head}

    for cur != nil {
        i := len(prev) - 1
        
        for cur.Val > prev[i].Val {
            prev[i - 1].Next = cur
            i--
        }
        prev = append(prev, cur)
        
        cur = cur.Next
    }

    return z.Next
}