/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    arr := []*ListNode{}
    cur := head
    
    for cur != nil {
        if len(arr) != 0 && cur.Val < arr[len(arr) - 1].Val {
            counter := 1
            
            for len(arr) != 0 && cur.Val < arr[len(arr) - 1].Val {
                cur.Val, arr[len(arr) - counter].Val = arr[len(arr) - counter].Val, cur.Val
                cur = arr[len(arr) - counter]
            }
            cur = arr[len(arr) - counter]
            arr = arr[:len(arr) - counter]
        } else {
            arr = append(arr, cur)
            cur = cur.Next
        }
    }
    return head
}