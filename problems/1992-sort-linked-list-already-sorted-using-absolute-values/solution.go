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
    
    for cur != nil {
        if cur.Val < 0 {
            prev.Next = prev.Next.Next
            cur.Next = head
            head = cur
            cur = prev.Next
        } else {
            prev = prev.Next
            cur = cur.Next
        }
    }
    
    return head
}