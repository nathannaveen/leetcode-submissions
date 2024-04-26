/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func frequenciesOfElements(head *ListNode) *ListNode {
    m := map[int] int {}

    cur := head

    for cur != nil {
        m[cur.Val]++
        cur = cur.Next
    }

    res := &ListNode{}
    cur2 := res
    prev := res

    for _, b := range m {
        cur2.Val = b
        cur2.Next = &ListNode{}
        prev = cur2
        cur2 = cur2.Next
    }

    prev.Next = nil

    return res
}
