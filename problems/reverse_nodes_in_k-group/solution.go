/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    counter := 0
    cur := head
    prev := head
    arr := make([]int, k)
    
    for cur != nil {
        if counter == 0 {
            prev = cur
        }
        arr[counter] = cur.Val
        counter++
        if counter == k {
            for i := 0; i < k; i++ {
                prev.Val = arr[len(arr) - 1 - i]
                prev = prev.Next
            }
            
            counter = 0
            arr = make([]int, k)
        }
        cur = cur.Next
    }
    
    return head
}