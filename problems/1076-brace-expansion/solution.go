func expand(s string) []string {
    options := [][]string{}
    open := false
    
    for _, letter := range s {
        if letter == '{' {
            open = true
            options = append(options, []string{})
        } else if letter == '}' {
            open = false
        } else if letter != ',' {
            if open {
                options[len(options) - 1] = append(options[len(options) - 1], string(letter))
            } else {
                options = append(options, []string{ string(letter) })
            }
        }
    }
    
    for i := 0; i < len(options); i++ {
        sort.Strings(options[i])
    }
    
    res := helper(options, 0, "")
    
    return res
}

func helper(options [][]string, i int, s string) []string {
    if i == len(options) {
        return []string{ s }
    }
    res := []string{}
    
    for _, letter := range options[i] {
        res = append(res, helper(options, i + 1, s + string(letter))...)
    }
    
    return res
}