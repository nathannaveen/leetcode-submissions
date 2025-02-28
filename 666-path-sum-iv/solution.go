type key struct {
    row, col int
}

func pathSum(nums []int) int {
    m := map[key] int {} // (row, col) : val

    for _, num := range nums {
        arr := []int{}
        for num > 0 {
            arr = append(arr, num % 10)
            num /= 10
        }
        m[key{arr[2], arr[1]}] = arr[0]
    }
    
    res, _ := dfs(1, 1, m)
    return res
}

func dfs(row, col int, m map[key] int) (int, int) {
    // return res, number of leaves
    if _, ok := m[key{row, col}]; !ok {
        return 0, 0
    }

    res := 0

    a, b := dfs(row + 1, col * 2, m)
    c, d := dfs(row + 1, col * 2 - 1, m)

    multiply := int(math.Max(float64(b + d), float64(1)))
    res += m[key{row, col}] * multiply + a + c

    return res, multiply
}