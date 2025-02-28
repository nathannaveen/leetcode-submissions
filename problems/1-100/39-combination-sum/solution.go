func combinationSum(candidates []int, target int) [][]int {
    return helper(0, candidates, target, []int{}, 0)
}

func helper(index int, candidates []int, target int, path []int, sum int) [][]int {
    if sum > target { return [][]int{} }
    if sum == target {
        return [][]int{ append([]int{}, path...) }
    }
    
    res := [][]int{}
    
    for i := index; i < len(candidates); i++ {
        res = append(res, helper(i, candidates, target, append(path, candidates[i]), sum + candidates[i])...)
    }
    return res
}