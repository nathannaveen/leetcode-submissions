/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    cur := head
    cur2 := head
    length := 0
    
    for cur != nil && cur.Next != nil {
        cur = cur.Next.Next
        length += 2
    }
    
    if cur != nil { length++ }
    
    temp := length - n - 1
    
    if temp == -1 { return cur2.Next }
    
    if temp < 0 { return nil }
    
    for i := 0; i < temp; i++ { cur2 = cur2.Next }
    
    next := cur2.Next.Next
    cur2.Next = next
    
    return head
}