func rowAndMaximumOnes(mat [][]int) []int {
    res := []int{0,0}

    for i, row := range mat {
        sum := 0
        for _, cell := range row {
            sum += cell
        }

        if sum > res[1] {
            res = []int{i, sum}
        }
    }

    return res
}