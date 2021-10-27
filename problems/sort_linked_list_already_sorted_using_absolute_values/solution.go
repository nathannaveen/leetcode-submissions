/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortLinkedList(head *ListNode) *ListNode {
    cur := head.Next
    prev := head
    start := head
    
    for cur != nil {
        if cur.Val < 0 {
            start = &ListNode{ cur.Val, start }
            cur = cur.Next
            prev.Next = prev.Next.Next
        } else {
            prev = prev.Next
            cur = cur.Next
        }
    }
    
    return start
}