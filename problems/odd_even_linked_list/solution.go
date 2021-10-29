/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil { return head }
    
    even := head
    oddHead := head.Next
    isOdd := true
    cur := head.Next
    prev := head

    for cur != nil {
        if !isOdd {
            temp := cur.Val
            cur = cur.Next
            prev.Next = cur
            temp2 := &ListNode{ temp, oddHead }
            even.Next = temp2
            even = even.Next
        } else {
            cur = cur.Next
            prev = prev.Next
        }

        isOdd = !isOdd
    }
    
    return head
}