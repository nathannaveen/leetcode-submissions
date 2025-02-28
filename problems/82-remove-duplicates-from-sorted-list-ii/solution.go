/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    prev := &ListNode{-101, head}
    res := prev
    
    for prev.Next != nil {
        val := prev.Next.Val
        change := false
        
        for prev.Next.Next != nil && prev.Next.Next.Val == val {
            prev.Next = prev.Next.Next
            change = true
        }
        
        if change {
            prev.Next = prev.Next.Next
        } else {
            prev = prev.Next
        }
    }
    return res.Next
}