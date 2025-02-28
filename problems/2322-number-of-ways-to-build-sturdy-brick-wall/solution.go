type key struct {
    preRow int
    height int
}

var dp = map[key] int {}

func buildWall(height int, width int, bricks []int) int {
    dp = map[key] int {}
    rows := findRows(bricks, width, 0, 0)

    return helper(rows, 0, height)
}

func findRows(bricks []int, width, curWidth int, location int) map[int] bool {
    res := map[int] bool {}
    
    for i := 0; i < len(bricks); i++ {
        if curWidth + bricks[i] == width {
            res[location] = true
        } else if curWidth + bricks[i] < width {
            x := findRows(bricks, width, curWidth + bricks[i], location | (1 << (curWidth + bricks[i] - 1)))

            for k := range x {
                res[k] = true
            }
        }
    }

    return res
}

func helper(rows map[int] bool, prevRow, height int) int {
    if height == 0 {
        return 1
    }

    if val, ok := dp[key{prevRow, height}]; ok {
        return val
    }

    res := 0

    for row := range rows {
        if row & prevRow == 0 {
            res += helper(rows, row, height - 1)
        }
    }

    res %= 1000000007

    dp[key{prevRow, height}] = res

    return res
}