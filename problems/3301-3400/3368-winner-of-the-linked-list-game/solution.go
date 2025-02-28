/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func gameResult(head *ListNode) string {
    cur := head
    e, o := 0, 0

    for cur != nil {
        if cur.Val > cur.Next.Val {
            e++
        } else {
            o++
        }
        cur = cur.Next.Next
    }

    if e > o {
        return "Even"
    } else if o > e {
        return "Odd"
    } else {
        return "Tie"
    }
}
