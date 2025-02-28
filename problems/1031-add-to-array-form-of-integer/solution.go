func addToArrayForm(num []int, k int) []int {
    carry := 0
    i := len(num) - 1
    res := []int{}

    for i >= 0 || k > 0 {
        n := carry
        if i >= 0 {
            n += num[i]
            i--
        }
        n += k % 10
        k /= 10

        if n > 9 {
            carry = 1
            n -= 10
        } else {
            carry = 0
        }

        res = append([]int{n}, res...)
    }
    if carry != 0 {
        res = append([]int{1}, res...)
    }

    return res
}