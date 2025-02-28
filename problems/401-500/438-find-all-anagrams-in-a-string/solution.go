func findAnagrams(s string, p string) []int {
    letters := make(map[byte] int)
    res := []int{}
    right, left := 0, 0
    
    for i := range p {
        letters[p[i]]++
    }
    
    for right < len(s) {
        if letters[s[right]] > 0 {
            letters[s[right]]--
            right++
        } else if letters[s[right]] == 0 {
            letters[s[left]]++
            left++
        } else if left == right {
            right++
            left++
        }
        
        if right - left == len(p) {
            res = append(res, left)
        }
    }
    
    return res
}