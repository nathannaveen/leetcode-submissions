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
    root.Val = 0
    m := helper(root)
    m[0] = true
    return FindElements{m}
}

func helper(root *TreeNode) map[int] bool {
    m := map[int] bool {}

    if root.Left != nil {
        v := root.Val * 2 + 1
        m[v] = true
        root.Left.Val = v
        m2 := helper(root.Left)
        for k, v := range m2 {
            m[k] = v
        }
    }
    if root.Right != nil {
        v := root.Val * 2 + 2
        m[v] = true
        root.Right.Val = v
        m2 := helper(root.Right)
        for k, v := range m2 {
            m[k] = v
        }
    }

    return m
}


func (this *FindElements) Find(target int) bool {
    return this.m[target]
}


/**
 * Your FindElements object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Find(target);
 */
