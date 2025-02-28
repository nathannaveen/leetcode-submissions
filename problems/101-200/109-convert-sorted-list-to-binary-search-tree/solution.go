/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    
    prev, turtle, hare := head, head, head

    for hare != nil && hare.Next != nil {
        hare = hare.Next.Next
        prev = turtle
        turtle = turtle.Next
    }

    res := &TreeNode{
        Val : turtle.Val,
    }

    if turtle.Next != nil {
        res.Right = sortedListToBST(turtle.Next)
    } 

    if prev.Next != nil {
        prev.Next = nil
        res.Left = sortedListToBST(head)
    }
    return res
}