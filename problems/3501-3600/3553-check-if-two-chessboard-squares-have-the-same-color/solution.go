func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
    l1, n1 := coordinate1[0], coordinate1[1]
    l2, n2 := coordinate2[0], coordinate2[1]
    
    return helper(l1, n1) == helper(l2, n2)
}

func helper(l, n byte) bool {
    x, _ := strconv.Atoi(string(n))
    if int(l - 'a') % 2 == 0 {
        if x % 2 == 0 {
            return true
        }
    } else {
        if x % 2 == 1 {
            return true
        }
    }

    return false
}
