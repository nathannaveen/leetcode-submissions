/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    res := &ListNode{}
    cur := res
    done := false
    
    for !done {
        done = true
        min := 10000
        
        for _, list := range lists {
            if list != nil && list.Val < min {
                min = list.Val
                done = false
            }
        }
        for i := range lists {
            if lists[i] != nil && lists[i].Val == min {
                cur.Next = &ListNode{ Val: min }
                cur = cur.Next
                lists[i] = lists[i].Next
            }
        }
    }
    
    return res.Next
}