func smallestSubsequence(s string) string {
    q := make([]byte, 0)
    seen := make(map[byte]bool)
    for i, _ := range s {
        for len(q)>0 && q[len(q)-1]>s[i] && !seen[s[i]] && strings.Contains(s[i+1:], string(q[len(q)-1])) {
            seen[q[len(q)-1]] = false
            q = q[:len(q)-1]
        }
        if !seen[s[i]] {
            q = append(q, s[i])
            seen[s[i]] = true
        }
        
    }
    return string(q)
}