/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    m := make(map[*ListNode] bool)
    
    cur := head
    
    for cur != nil {
        if m[cur] {
            return cur
        }
        m[cur] = true
        
        cur = cur.Next
    }
    return nil
}