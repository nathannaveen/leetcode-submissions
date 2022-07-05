/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {
    res := make([][]int, m)
    
    for i := 0; i < m; i++ {
        res[i] = make([]int, n)
        
        for j := 0; j < n; j++ {
            res[i][j] = -1
        }
    }
    
    directions := [][]int{ {0, 1}, {1, 0}, {0, -1}, {-1, 0} }
    dCounter := 0
    
    cur := head
    
    i, j := 0, 0
    
    for cur != nil {
        res[i][j] = cur.Val
        
        I := i + directions[dCounter % len(directions)][0]
        J := j + directions[dCounter % len(directions)][1]
        
        if I < 0 || J < 0 || I >= m || J >= n || res[I][J] != -1 {
            dCounter++
        }
        
        i += directions[dCounter % len(directions)][0]
        j += directions[dCounter % len(directions)][1]
        
        cur = cur.Next
    }
    
    return res
}