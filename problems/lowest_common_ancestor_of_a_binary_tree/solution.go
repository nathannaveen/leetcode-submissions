/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type key struct {
    arr []*TreeNode
    node *TreeNode
}

var Q = &TreeNode{}
var qKey = key{}
var P = &TreeNode{}
var pKey = key{}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    P = p
    pKey = key{}
    Q = q
    qKey = key{}
    helper(key{ []*TreeNode{}, root })
    pMap := make(map[*TreeNode] bool)
    
    for _, node := range pKey.arr {
        pMap[node] = true
    }
    
    for i := len(qKey.arr) - 1; i >= 0; i-- {
        if pMap[qKey.arr[i]] {
            return qKey.arr[i]
        }
    }
    
    return root
}

func helper(k key) {
    if k.node != nil {
        k.arr = append(k.arr, k.node)
        if k.node == P {
            pKey = k
        } else if k.node == Q {
            qKey = k
        }
        
        helper(key{ k.arr, k.node.Right })
        helper(key{ k.arr, k.node.Left })
    }
}