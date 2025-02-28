/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type key struct {
    num int
    lenDashes int
}

type key2 struct {
    node *TreeNode
    pos int
}

func recoverFromPreorder(S string) *TreeNode {
    arr := []key{}
    
    d := 0
    s := ""
    
    for _, i := range S {
        if i == '-' && s != "" {
            temp, _ := strconv.Atoi(s)
            arr = append(arr, key{ temp, d })
            s = ""
            d = 0
        } 
        if i == '-' {
            d++
        } else if int(i) <= 57 && int(i) >= 48 {
            s += string(i)
        }
    }
    temp, _ := strconv.Atoi(s)
    arr = append(arr, key{ temp, d })
    
    res := &TreeNode{Val: arr[0].num, }
    tree := res
    
    stack := []key2{ key2{tree, 0} }
    
    for i := 1; i < len(arr); i++ {
        for true {
            if stack[len(stack) - 1].pos == arr[i].lenDashes - 1 && 
            (stack[len(stack) - 1].node.Left == nil || stack[len(stack) - 1].node.Right == nil) {
                break
            }
            stack = stack[:len(stack) - 1]
        }
        
        if stack[len(stack) - 1].node.Left != nil {
            stack[len(stack) - 1].node.Right = &TreeNode{Val: arr[i].num,}
            stack = append(stack, key2{ stack[len(stack) - 1].node.Right, stack[len(stack) - 1].pos + 1 })
        } else {
            stack[len(stack) - 1].node.Left = &TreeNode{Val: arr[i].num,}
            stack = append(stack, key2{ stack[len(stack) - 1].node.Left, stack[len(stack) - 1].pos + 1 })
        }
    }
    
    return res
}