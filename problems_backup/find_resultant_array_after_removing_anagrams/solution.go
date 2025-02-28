func removeAnagrams(words []string) []string {
    res := []string{ words[0] }
    
    for i := 1; i < len(words); i++ {
        if !isAnagram(res[len(res) - 1], words[i]) {
            res = append(res, words[i])
        }
    }
    return res
}

func isAnagram(a, b string) bool {
    if len(a) != len(b) { 
        return false 
    }
    
    A, B := [26]byte{}, [26]byte{}
    
    for i := range a {
        A[a[i] - 'a']++
        B[b[i] - 'a']++
    }
    
    return A == B
}