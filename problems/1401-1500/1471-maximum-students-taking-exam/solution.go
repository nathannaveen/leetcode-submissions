type key struct {
    actualSeats int
    nextColCantBe int
    numberOccupied int
}

type key2 struct {
    col, cantBe int
}

var dp = map[key2] int {}

func maxStudents(seats [][]byte) int {
    dp = map[key2] int {}
    keysPerCol := map[int] []key {}

    for j := 0; j < len(seats[0]); j++ {
        keysPerCol[j] = allForACol(seats, 0, j)
    }

    return dfs(0, keysPerCol, 0)
}

func dfs(col int, keysPerCol map[int] []key, colCantBe int) int {
    res := 0

    if val, ok := dp[key2{col, colCantBe}]; ok {
        return val
    }

    for _, key := range keysPerCol[col] {
        if key.actualSeats & colCantBe == 0 {
            res = max(key.numberOccupied +  dfs(col + 1, keysPerCol, key.nextColCantBe), res)
        }
    }

    dp[key2{col, colCantBe}] = res

    return res
}

func allForACol(seats [][]byte, row, col int) []key {
    if row == len(seats) {
        return []key{ {0, 0, 0} }
    }

    res := []key{}

    arr := allForACol(seats, row + 1, col)

    if seats[row][col] == '.' {
        for _, a := range arr {
            actual := a.actualSeats | (1 << row)
            nextColCantBe := a.nextColCantBe | (1 << row)
            if row != 0 {
                nextColCantBe |= (1 << (row - 1))
            }

            if row != len(seats) - 1 {
                nextColCantBe |= (1 << (row + 1))
            }

            res = append(res, key{actual, nextColCantBe, a.numberOccupied + 1})
        }
    }

    res = append(res, arr...)

    return res
}