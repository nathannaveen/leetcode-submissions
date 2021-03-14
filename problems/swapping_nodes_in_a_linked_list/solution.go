/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	g := head
	c := head
	l := head
	counter := 0

	for l != nil {
		counter++
		l = l.Next
	}

	for i := 0; i < counter-k; i++ {
		g = g.Next
	}

	for i := 0; i < k-1; i++ {
		c = c.Next
	}

	g.Val, c.Val = c.Val, g.Val
	return head
}