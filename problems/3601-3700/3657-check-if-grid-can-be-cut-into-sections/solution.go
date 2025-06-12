func checkValidCuts(n int, rec [][]int) bool {
    x := [][]int{}
    y := [][]int{}

    for _, r := range rec {
        x = append(x, []int{r[0], r[2]})
        y = append(y, []int{r[1], r[3]})
    }

    sort.Slice(x, func(i, j int) bool { return x[i][0] < x[j][0] })
    sort.Slice(y, func(i, j int) bool { return y[i][0] < y[j][0] })

    prevX := x[0]
    prevY := y[0]
    xCnt := 1
    yCnt := 1

    for i := 1; i < len(x); i++ {
        if x[i][0] < prevX[1] && x[i][0] >= prevX[0] {
            if x[i][1] > prevX[1] {
                prevX[1] = x[i][1]
            }
        } else {
            prevX = x[i]
            xCnt++
        }
    }

    for i := 1; i < len(y); i++ {
        if y[i][0] < prevY[1] && y[i][0] >= prevY[0] {
            if y[i][1] > prevY[1] {
                prevY[1] = y[i][1]
            }
        } else {
            prevY = y[i]
            yCnt++
        }
    }

    return xCnt >= 3 || yCnt >= 3
}
