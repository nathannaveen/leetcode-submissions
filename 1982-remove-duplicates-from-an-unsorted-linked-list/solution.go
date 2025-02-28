/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
    m := make(map[int] int)
    cur := head
    
    for cur != nil {
        m[cur.Val]++
        cur = cur.Next
    }
    newPrev := &ListNode{ Val: 0, Next: head}
    prev := newPrev
    for prev.Next != nil {
        if m[prev.Next.Val] > 1 {
            prev.Next = prev.Next.Next
        } else {
            prev = prev.Next
        }
    }
    return newPrev.Next
}