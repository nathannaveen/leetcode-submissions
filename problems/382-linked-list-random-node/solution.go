/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
    arr []int
}


/** @param head The linked list's head.
        Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
    arr := []int{}
    for head != nil {
        arr = append(arr, head.Val)
        head = head.Next
    }
    return Solution{ arr }
}


/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
    return this.arr[rand.Intn(len(this.arr))]
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */