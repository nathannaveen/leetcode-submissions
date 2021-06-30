/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, nums []int) int {
    m := make(map[int] int)
    res := 0
    prev := head
    cur := head
    
    for _, num := range nums {
        m[num] = 1
    }
    
    for cur != nil {
        if m[cur.Val] == 1 && (cur == head || m[prev.Val] == 0) {
            res++
        }
        prev = cur
        cur = cur.Next
    }
    
    return res
}