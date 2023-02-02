/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type key struct {
     val int
     index int
 }

func minimumOperations(root *TreeNode) int {
    // use bfs to find each level

    queue := []*TreeNode{root}
    cnt := 0

    for len(queue) != 0 {
        n := len(queue)
        arr := []key{}

        for i := 0; i < n; i++ {
            pop := queue[0]
            queue = queue[1:]
            arr = append(arr, key{pop.Val, i})
            if pop.Left != nil {
                queue = append(queue, pop.Left)
            }
            if pop.Right != nil {
                queue = append(queue, pop.Right)
            }
        }
        
        arr2 := make([]int, len(arr))
        
        for i, a := range arr { arr2[i] = a.val }

        sort.Slice(arr, func(i, j int) bool { 
            return arr[i].val < arr[j].val 
        })

        m := map[int] int {} // value : index

        for i, a := range arr { m[a.val] = i }

        for i := 0; i < len(arr2); i++ {
            if arr2[i] == arr[i].val {
                continue
            }
            arr2[i], arr2[arr[i].index], arr[i].index, arr[m[arr2[i]]].index = 
            arr2[arr[i].index], arr2[i], arr[m[arr2[i]]].index, arr[i].index
            cnt++
        }
    }

    return cnt
}












