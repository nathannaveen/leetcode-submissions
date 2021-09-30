func findSmallestRegion(regions [][]string, region1 string, region2 string) string {
    m := make(map[string] string)
    m2 := make(map[string] bool)
    
    for _, region := range regions {
        for i := 1; i < len(region); i++ {
            m[region[i]] = region[0]
        }
    }
    
    for m[region1] != "" {
        m2[region1] = true
        region1 = m[region1]
    }
    m2[region1] = true
    
    for m[region2] != "" {
        if m2[region2] { return region2 }
        region2 = m[region2]
    }
    
    return region2
}
