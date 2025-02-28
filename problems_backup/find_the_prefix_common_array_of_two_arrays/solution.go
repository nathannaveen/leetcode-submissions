func findThePrefixCommonArray(A []int, B []int) []int {
    m := map[int] int {}
    c := 0
    res := []int{}

    for i := 0; i < len(A); i++ {
        m[A[i]]++
        m[B[i]]--

        if m[A[i]] == 0 {
            c++
        }

        if A[i] != B[i] && m[B[i]] == 0 {
            c++
        }

        res = append(res, c)
    }

    return res
}