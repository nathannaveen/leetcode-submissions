func compareVersion(version1 string, version2 string) int {
    v1 := strings.Split(version1, ".")
    v2 := strings.Split(version2, ".")
    i := 0
    
    for i < len(v1) && i < len(v2) {
        val1, _ := strconv.Atoi(v1[i])
        val2, _ := strconv.Atoi(v2[i])
        
        if val1 < val2 { return -1 }
        if val1 > val2 { return 1 }
        
        i++
    }
    
    for j := i; j < len(v1); j++ {
        val, _ := strconv.Atoi(v1[j])
        if val != 0 { return 1 }
    }
    
    for j := i; j < len(v2); j++ {
        val, _ := strconv.Atoi(v2[j])
        if val != 0 { return -1 }
    }
    
    return 0
}