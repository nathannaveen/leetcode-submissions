/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    fast, slow := head, head
    max := 0
    arr := []int{}
    i := 0
    
    for slow != nil {
        if fast != nil {
            arr = append(arr, slow.Val)
            fast = fast.Next.Next
        } else {
            temp := arr[len(arr) - 1 - i] + slow.Val
            if temp > max {
                max = temp
            }
            i++
        }
        
        slow = slow.Next
    }
    return max
}