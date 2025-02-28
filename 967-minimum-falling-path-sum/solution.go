type pos struct {
    i, j int
}

var dp map[pos] int

func minFallingPathSum(matrix [][]int) int {
    dp = map[pos] int {}

    min := 10001

    for j := 0; j < len(matrix[0]); j++ {
        x := dfs(0, j, matrix)

        if x < min {
            min = x
        }
    }

    return min
}

func dfs(i, j int, matrix [][]int) int {
    if i >= len(matrix) || j >= len(matrix[0]) || j < 0  {
        return 10001
    }

    if val, ok := dp[pos{i, j}]; ok {
        return val
    }

    min := 10001

    for _, d := range []pos{ {1, -1}, {1, 0}, {1, 1} } {
        x := dfs(i + d.i, j + d.j, matrix)

        if x < min {
            min = x
        }
    }

    if min == 10001 {
        min = 0
    }

    dp[pos{i, j}] = min + matrix[i][j]

    return min + matrix[i][j]
}