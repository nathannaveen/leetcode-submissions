/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    arr := []*ListNode{}
    cur := head
    
    for cur != nil {
        if len(arr) != 0 && cur.Val < arr[len(arr) - 1].Val {
            cur.Val, arr[len(arr) - 1].Val = arr[len(arr) - 1].Val, cur.Val
            cur = arr[len(arr) - 1]
            arr = arr[:len(arr) - 1]
        } else {
            arr = append(arr, cur)
            cur = cur.Next
        }
    }
    return head
}