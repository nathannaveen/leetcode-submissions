/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    length := 0
    end := head
    cur := head
    
    for end != nil && end.Next != nil {
        length++
        end = end.Next
    }
    
    length++
    
    for i := 0; i < length - (k % length); i++ {
        if end != nil {
            end.Next = &ListNode{ Val: cur.Val, }
            end = end.Next
            cur = cur.Next
        }
    }
    
    return cur
}