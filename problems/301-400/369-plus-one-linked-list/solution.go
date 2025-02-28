/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var m map[*ListNode] *ListNode = make(map[*ListNode] *ListNode)

func plusOne(head *ListNode) *ListNode {
    start := &ListNode{0, head}
    cur := head 
    prev := start
    carry := 1
    m = make(map[*ListNode] *ListNode)
    
    for cur.Next != nil {
        m[cur] = prev
        
        prev = cur
        cur = cur.Next
    }
    
    m[cur] = prev
    
    for cur.Val + carry >= 10 {
        cur.Val = 0
        cur = m[cur]
        carry = 1
    }
    
    cur.Val += carry
    
    if start.Val == 0 {
        return head
    }
    
    return start
}
