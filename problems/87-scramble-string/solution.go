var dp = map[[2]string] bool {}

func isScramble(s1 string, s2 string) bool {
    dp = map[[2]string] bool {}
    return helper(s1, s2)
}

func helper(s1, s2 string) bool {
    if s1 == s2 || len(s1) == 1 {
        return s1 == s2
    }

    if val, ok := dp[[2]string{s1, s2}]; ok {
        return val
    }

    for i := 1; i < len(s1); i++ {
        x := (helper(s1[:i], s2[:i]) && helper(s1[i:], s2[i:])) || 
        (helper(s1[i:], s2[:len(s2) - i]) && helper(s1[:i], s2[len(s2) - i:]))

        if x {
            dp[[2]string{s1, s2}] = true
            return true
        }
    }

    dp[[2]string{s1, s2}] = false
    return false
}