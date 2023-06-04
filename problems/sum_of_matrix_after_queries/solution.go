type key struct {
    v, t int
}

func matrixSumQueries(n int, queries [][]int) int64 {
    c, r := make([]key, n), make([]key, n)
    res := int64(0)
    
    for i, q := range queries {
        if q[0] == 0 {
            r[q[1]] = key{q[2], i + 1}
        } else {
            c[q[1]] = key{q[2], i + 1}
        }
    }
    
    for _, i := range r {
        for _, j := range c {
            if i.t < j.t {
                res += int64(j.v)
            } else {
                res += int64(i.v)
            }
        }
    }
    
    return res
}