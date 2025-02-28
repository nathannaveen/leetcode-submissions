/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    counter := 0
    cur := list1
    start := list1
    end := list2
    
    for end.Next != nil {
        end = end.Next
    }
    
    for cur != nil {
        if counter == a - 1 {
            start = cur
        } else if counter == b + 1 {
            start.Next = list2
            end.Next = cur
            break
        }
        counter++
        cur = cur.Next
    }
    
    return list1
}