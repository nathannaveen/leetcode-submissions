/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    prev := head
    cur := head
    
    for cur != nil {
        if cur.Val != 0 {
            prev.Val += cur.Val
            cur = prev
            cur.Next = cur.Next.Next
        } else {
            if cur.Next == nil {
                prev.Next = nil
            }
            prev = cur
        }
        cur = cur.Next
    }
    
    return head
}