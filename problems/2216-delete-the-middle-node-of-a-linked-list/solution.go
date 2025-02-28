/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head.Next == nil { return nil }
    
    hare, prev := head, &ListNode{ -1, head }
    
    for hare != nil && hare.Next != nil {
        hare = hare.Next.Next
        prev = prev.Next
    }
    
    prev.Next = prev.Next.Next
    
    return head
}