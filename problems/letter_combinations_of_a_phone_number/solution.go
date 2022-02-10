func letterCombinations(digits string) []string {
    if len(digits) == 0{
        return []string{}
    }
    
    return helper(digits, "")
}

func helper(digits, s string) []string {
    res := []string{}
    if len(digits) == 0 { 
        return []string{ s }
    }
    
    arr := []string{ 
        "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz", 
    }
    
    for _, letter := range arr[digits[0] - '2'] {
        res = append(res, helper(digits[1:], s + string(letter))...)
    }
    return res
}