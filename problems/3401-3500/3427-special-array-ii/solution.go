func isArraySpecial(nums []int, queries [][]int) []bool {
    lastPar := make([]int, len(nums)) // index of latest parity
    prev := -1
    last := -1
    for i, n := range nums {
        if prev != -1 && n % 2 == prev % 2 {
            last = i
        }
        prev = n
        lastPar[i] = last
    }

    res := make([]bool, len(queries))

    for i, query := range queries {
        if lastPar[query[0]] == lastPar[query[1]] {
            res[i] = true
        } else {
            res[i] = false
        }
    }

    return res
}
