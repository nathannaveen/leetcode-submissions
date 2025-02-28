func numberOfWays(s string) int64 {
    m := map[string] int64 { "0" : 0, "1" : 0, "01" : 0, "10" : 0, "010" : 0, "101" : 0 }
    
    for _, l := range s {
        letter := string(l)
        
        if letter == "0" {
            m["0"]++
            m["10"] += m["1"]
            m["010"] += m["01"]
        } else {
            m["1"]++
            m["01"] += m["0"]
            m["101"] += m["10"]
        }
    }
    
    return m["101"] + m["010"]
}