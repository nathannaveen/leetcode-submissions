func maximumSetSize(nums1 []int, nums2 []int) int {
    one := map[int] bool {}
    two := map[int] bool {}
    nodes := []int{}
    nodesMap := map[int] bool {}
    
    for _, n := range nums1 {
        one[n] = true
        if !nodesMap[n] {
            nodes = append(nodes, n)
            nodesMap[n] = true
        }
    }
    
    for _, n := range nums2 {
        two[n] = true
        if !nodesMap[n] {
            nodes = append(nodes, n)
            nodesMap[n] = true
        }
    }
    
    i := 0
    
    for len(one) > len(nums1) / 2 {
        if i == len(nodes) {
            break
        }
        
        if one[nodes[i]] && two[nodes[i]] {
            delete(one, nodes[i])
        }
        i++
    }
    
    i = 0
    for len(two) > len(nums2) / 2 {
        if i == len(nodes) {
            break
        }
        
        if one[nodes[i]] && two[nodes[i]] {
            delete(two, nodes[i])
        }
        i++
    }
    
    overlap := map[int] bool {}
    
    cnt := 0
    for k := range one {
        if cnt == len(nums1) / 2 {
            break
        }
        cnt++
        overlap[k] = true
    }
    
    cnt = 0
    for k := range two {
        if cnt == len(nums2) / 2 {
            break
        }
        cnt++
        overlap[k] = true
    }
    
    return len(overlap)
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}