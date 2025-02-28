func minMoves(target int, maxDoubles int) int {
    res := 0
    for target > 1 && maxDoubles > 0 {
        if target % 2 == 1 {
            target--
        } else if maxDoubles > 0 {
            maxDoubles--
            target /= 2
        } else {
            target--
        }
        res++
    }
    return res + target - 1
}