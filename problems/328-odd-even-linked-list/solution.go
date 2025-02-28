/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil { return head }
    
    evenPos, oddHead := head, head.Next
    isEven := false
    cur, prev := head.Next, head

    for cur != nil {
        if isEven {
            temp := cur.Val
            cur = cur.Next
            prev.Next = cur
            newNode := &ListNode{ temp, oddHead }
            evenPos.Next = newNode
            evenPos = evenPos.Next
        } else {
            cur = cur.Next
            prev = prev.Next
        }

        isEven = !isEven
    }
    
    return head
}