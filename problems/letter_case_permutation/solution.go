func letterCasePermutation(s string) []string {
    return helper(s, "")
}

func helper(s string, a string) []string {
    if len(s) == 0 {
        return []string{ a }
    }
    
    res := []string{}
    letter := s[0] // first letter
    s = s[1:]
    
    if letter >= 'a' { // lowercase
        res = append(res, helper(s, a + string(letter - 32))...)
    } else if letter >= 'A' { // uppercase
        res = append(res, helper(s, a + string(letter + 32))...)
    }
    res = append(res, helper(s, a + string(letter))...)
    
    return res
}