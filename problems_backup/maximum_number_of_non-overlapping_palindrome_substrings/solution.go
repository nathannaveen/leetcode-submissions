func maxPalindromes(s string, k int) int {
    res := 0

    for i := 0; i < len(s); i++ {
        for j := i; j < len(s); j++ {
            l := j - i + 1
            if l > k + 1 {
                break
            }

            if l >= k && isPalindrome(s, i, j) {
                res++
                i = j
                break
            }
        }
    }

    return res
}

func isPalindrome(s string, i, j int) bool {
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}