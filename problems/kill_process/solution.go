func killProcess(pid []int, ppid []int, kill int) []int {
    m := make(map[int] []int )
	
    for i := 0; i < len(pid); i++ {
        m[ppid[i]] = append(m[ppid[i]], pid[i])
    }
    
    return helper(kill, m, []int{})
}

func helper(kill int, m map[int] []int, res []int) []int {
    res = append(res, kill)
    
    for _, n := range m[kill] {
        res = helper(n, m, res)
    }
    
    return res
}