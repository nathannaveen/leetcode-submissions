/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
    prev, cur, next := head, head.Next, head.Next.Next
    prevIndex, firstIndex := -1, -1
    index, counter := 1, 0
    res := []int{100000, 0}
    
    
    for next != nil {
        if (cur.Val < prev.Val && cur.Val < next.Val) || 
        (cur.Val > prev.Val && cur.Val > next.Val) {
            if prevIndex != -1 && index - prevIndex < res[0] {
                res[0] = index - prevIndex
            } 
            
            if prevIndex != -1 {
                res[1] = index - firstIndex
            } else {
                firstIndex = index
            }
            prevIndex = index
            counter++
        }
        
        index++
        prev, cur, next = prev.Next, cur.Next, next.Next
    }
    
    if counter < 2 { return []int{ -1, -1 } }

    return res
}