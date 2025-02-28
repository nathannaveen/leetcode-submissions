var res = [][]int{}

func combine(n int, k int) [][]int {
    res = [][]int{}
    helper(k, n, 0, []int{})
    return res
}

func helper(k, n, cur int, arr []int) {
    if k == 0 {
        res = append(res, append([]int{}, arr...))
        return
    }
    
    if cur == n {
        return
    }
    
    for i := cur + 1; i <= n; i++ {
        helper(k - 1, n, i, append(arr, i))
    }
}