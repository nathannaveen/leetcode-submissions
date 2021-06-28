/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    cur := head
    
    for cur != nil && cur.Next != nil {
        cur.Val, cur.Next.Val = cur.Next.Val, cur.Val
        cur = cur.Next.Next
    }
    return head
}