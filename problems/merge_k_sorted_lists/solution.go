/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    arr := []int{}
    for _, head := range lists {
        cur := head
        
        for cur != nil {
            arr = append(arr, cur.Val)
            cur = cur.Next
        }
    }
    
    sort.Ints(arr)
    
    if len(arr) > 0 {
        res := &ListNode{
            Val: arr[0],
        }
        cur := res
        for i := 1; i < len(arr); i++ {
            cur.Next = &ListNode{
                Val: arr[i],
            }
            cur = cur.Next
        }
        return res
    }
    return nil
}