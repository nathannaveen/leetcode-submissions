/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    arr := []int{}
    
    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }
    
    left, right := 0, len(arr) - 1
    
    for left < right {
        if arr[left] != arr[right] { return false }
        
        left++
        right--
    }
    
    return true
}