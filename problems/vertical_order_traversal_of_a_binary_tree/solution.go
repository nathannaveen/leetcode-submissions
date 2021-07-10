type key struct {
    val int
    row int
    col int
}

type key2 struct {
    node *TreeNode
    row int
    col int
}

type key3 struct {
    val int
    row int
}

func verticalTraversal(root *TreeNode) [][]int {
	res := [][]int{}
    arr := []key{ key{ root.Val, 0, 0 } }
    queue := []key2{ key2{ root, 0, 0 } }
    min := root.Val
    max := root.Val
    
    for len(queue) != 0 {
        n := len(queue)
        
        for i := 0; i < n; i++ {
            node := queue[0]
            queue = queue[1:]
            
            if node.node.Left != nil {
                arr = append(arr, key{ node.node.Left.Val, node.row + 1, node.col - 1 })
                queue = append(queue, key2{ node.node.Left, node.row + 1, node.col - 1 })
                min = int(math.Min(float64(node.col - 1), float64(min)))
            }
            if node.node.Right != nil {
                arr = append(arr, key{ node.node.Right.Val, node.row + 1, node.col + 1 })
                queue = append(queue, key2{ node.node.Right, node.row + 1, node.col + 1 })
                max = int(math.Max(float64(node.col + 1), float64(max)))
            }
        }
    }
    
    arr2 := make([][]key3, max - min + 1)
    
    for _, i := range arr {
        arr2[i.col - min] = append(arr2[i.col - min], key3{i.val, i.row})
    }
    
    for _, k := range arr2 {
        arr3 := []int{}
        sort.Slice(k, func(i, j int) bool {
            if k[i].row == k[j].row {
                return k[i].val < k[j].val
            }
            return k[i].row < k[j].row
        })
        
        for _, i := range k {
            arr3 = append(arr3, i.val)
        }
        res = append(res, arr3)
    }
    
    counter := 1
    
    for len(res[len(res) - counter]) == 0 {
        if len(res[len(res) - counter]) == 0 {
            counter++
            continue
        }
        break
    }
    
    return res[:len(res) - counter + 1]
}