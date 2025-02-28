func countKConstraintSubstrings(s string, k int) int {
    start := 0
    zeros, ones := 0, 0

    res := 0

    for i := 0; i < len(s); i++ {
        zeros, ones = conditionalAdd(s[i], 1, zeros, ones)

        for zeros > k && ones > k {
            zeros, ones = conditionalAdd(s[start], -1, zeros, ones)
            start++
        }
        
        n := i - start + 1
        res += n
    }

    return res
}

func conditionalAdd(n byte, add, zeros, ones int) (int, int) {
    if n == '0' {
        zeros += add
    } else {
        ones += add
    }
    return zeros, ones
}
