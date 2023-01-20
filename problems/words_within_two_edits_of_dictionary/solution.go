func twoEditWords(queries []string, dictionary []string) []string {
    res := []string{}

    for _, q := range queries {
        for _, d := range dictionary {
            diff := 0
            for i := 0; i < len(q); i++ {
                if q[i] != d[i] {
                    diff++
                }
            }
            if diff <= 2 {
                res = append(res, q)
                break
            }
        }
    }

    return res
}