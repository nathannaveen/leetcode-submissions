/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type FindElements struct {
    m map[int] bool
}

func Constructor(root *TreeNode) FindElements {
    m := map[int] bool {}
    makeTree(root, 0, m)
    return FindElements{ m }
}

func makeTree(root *TreeNode, val int, m map[int] bool) {
    if root == nil { return }
    root.Val = val
    m[val] = true
    makeTree(root.Left, val * 2 + 1, m)
    makeTree(root.Right, val * 2 + 2, m)
}


func (this *FindElements) Find(target int) bool {
    return this.m[target]
}


/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */