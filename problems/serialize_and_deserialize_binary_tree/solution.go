/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    s string
}

func Constructor() Codec {
    return Codec{""}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    queue := []*TreeNode{root}
    
    for len(queue) != 0 {
        n := len(queue)
        
        for i := 0; i < n; i++ {
            pop := queue[0]
            queue = queue[1:]
            
            if pop == nil {
                this.s += " nil"
            } else {
                this.s += " " + strconv.Itoa(pop.Val)
                queue = append(queue, pop.Left, pop.Right)
            }
        }
    }
    
    return this.s
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    arr := strings.Fields(data)
    if arr[0] == "nil" {
        return nil
    }
    h, _ := strconv.Atoi(arr[0])
    res := &TreeNode{Val: h,}
    queue := []*TreeNode{res}
    counter := 1
    
    for len(queue) != 0 {
        n := len(queue)
        
        for i := 0; i < n; i++ {
            pop := queue[0]
            queue = queue[1:]
            
            if counter < len(arr) && arr[counter] != "nil" {
                temp, _ := strconv.Atoi(arr[counter])
                pop.Left = &TreeNode{Val: temp,}
                queue = append(queue, pop.Left)
            }
            
            counter ++
            
            if counter < len(arr) && arr[counter] != "nil" {
                temp, _ := strconv.Atoi(arr[counter])
                pop.Right = &TreeNode{Val: temp,}
                queue = append(queue, pop.Right)
            }
            
            counter ++
        }
    }
    
    return res
}


/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */