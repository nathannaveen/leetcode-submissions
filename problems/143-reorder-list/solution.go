/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode)  {
    cur := head.Next
    prev := head
    arr := []*ListNode{} // prevValues
    
    for cur != nil {
        arr = append(arr, prev)
        cur = cur.Next
        prev = prev.Next
    }
    
    left, right := 0, len(arr) - 1
    
    for left < right {
        
        temp := arr[right].Next
        temp2 := arr[left].Next
        arr[right].Next = arr[right].Next.Next
        
        arr[left].Next = temp
        arr[left].Next.Next = temp2
        
        left++
        right--
    }
}