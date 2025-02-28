/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	num := 0

	for head != nil {
		num *= 2
		if head.Val == 1 {
			num += 1
		}
		head = head.Next
	}

	return num
}