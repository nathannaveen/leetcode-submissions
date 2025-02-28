/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    cur := head
    arr := make([]int, right)
    for i := 0; i < right; i++ {
        arr[i] = cur.Val
        cur = cur.Next
    }
    
    cur = head
    
    for i := 0; i < right; i++ {
        if i + 1 >= left {
            cur.Val = arr[len(arr) - ((i + 1) - left) - 1]
        }
        cur = cur.Next
    }
    
    return head
}