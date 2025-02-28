/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    start := &ListNode{Next: head}
    prev := start
    cur := head
    greaterThanXStart := &ListNode{}
    greaterThanX := greaterThanXStart
    
    for cur != nil {
        if cur.Val >= x {
            prev.Next = prev.Next.Next
            temp := cur
            greaterThanX.Next = temp
            greaterThanX = greaterThanX.Next
            cur = prev.Next
        } else {
            cur = cur.Next
            prev = prev.Next
        }
    }
    
    greaterThanX.Next = nil
    prev.Next = greaterThanXStart.Next
    return start.Next
}