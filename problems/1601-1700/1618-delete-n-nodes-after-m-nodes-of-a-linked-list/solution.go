/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNodes(head *ListNode, m int, n int) *ListNode {
    cur := head
    prev := head
    i := 1
    
    for cur != nil {
        
        if i == m {
            prev = cur
        }
        
        if i == n + m + 1  { 
            prev.Next = cur
            i = 1
            if cur.Next == nil {
                return head
            }
        } else {
            i++
            cur = cur.Next    
        }
        
    }
    if i > m {
        prev.Next = cur
    }
    return head
}