/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{}
    stack1, stack2 := []int{}, []int{}
    
    for l1 != nil || l2 != nil {
        if l1 != nil { 
            stack1 = append(stack1, l1.Val) 
            l1 = l1.Next
        }
        if l2 != nil { 
            stack2 = append(stack2, l2.Val) 
            l2 = l2.Next
        }
    }
    
    sum := 0
    
    for len(stack1) != 0 || len(stack2) != 0 {
        if len(stack1) != 0 {
            sum += stack1[len(stack1) - 1]
            stack1 = stack1[ :len(stack1) - 1]
        }
        if len(stack2) != 0 {
            sum += stack2[len(stack2) - 1]
            stack2 = stack2[ :len(stack2) - 1]
        }
        res.Val = sum % 10
        res = &ListNode{ sum / 10, res }
        sum /= 10
    }
    if res.Val == 0 { return res.Next }
    return res
}