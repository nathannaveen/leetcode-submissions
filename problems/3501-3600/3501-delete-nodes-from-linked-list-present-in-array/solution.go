/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func modifiedList(nums []int, head *ListNode) *ListNode {
    m := map[int] bool {}
    for _, num := range nums {
        m[num] = true
    }

    newHead := &ListNode{Val: 0, Next: head}
    cur := head
    prev := newHead

    for cur != nil {
        if !m[cur.Val] {
            prev.Next = cur
            prev = prev.Next
        }

        cur = cur.Next
    }
    prev.Next = nil
    return newHead.Next
}
