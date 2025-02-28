func canConstruct(ransomNote string, magazine string) bool {
    arr := make([]int, 26)
    
    for _, l := range magazine {
        arr[l - 'a']++
    }
    for _, l := range ransomNote {
        arr[l - 'a']--
        
        if arr[l - 'a'] < 0 {
            return false
        }
    }
    
    return true
}