/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    cur := head

    for cur.Next != nil {
        cur.Next = &ListNode{ 
            Val: GCD(cur.Val, cur.Next.Val),
            Next: cur.Next,
        }
        cur = cur.Next.Next
    }

    return head
}

func GCD(a, b int) int {
	for b > 0 {
        a, b = b, a % b
	}
	return a
}