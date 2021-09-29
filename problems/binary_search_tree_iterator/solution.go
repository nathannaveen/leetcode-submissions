/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    arr []int
    i int
}

func Constructor(root *TreeNode) BSTIterator {
    arr := helper(root)
    
    sort.Ints(arr)
    
    return BSTIterator{arr, 0}
}

func helper(root *TreeNode) []int {
    res := []int{}
    if root != nil {
        res = append(res, root.Val)
        res = append(res, helper(root.Right)...)
        res = append(res, helper(root.Left)...)
    }
    return res
}

func (this *BSTIterator) Next() int {
    this.i++
    return this.arr[this.i - 1]
}


func (this *BSTIterator) HasNext() bool {
    return this.i < len(this.arr)
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */