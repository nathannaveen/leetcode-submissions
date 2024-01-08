func maximumLength(s string) int {
    res := -1

    m := map[key] int {} // key{special letter, length of the substring} : number of this

    l := ""
    cnt := 0

    for i := 0; i < len(s); i++ {
        if string(s[i]) != l {
            cnt = 1
            l = string(s[i])
            m[key{l, cnt}]++

            if cnt > res && m[key{l, cnt}] >= 3 {
                res = cnt
            }
        } else {
            cnt++
            for j := 1; j <= cnt; j++ {
                m[key{l, j}]++

                if j > res && m[key{l, j}] >= 3 {
                    res = j
                }
            }
        }
    }
    
    return res
}

type key struct {
    l string
    n int
}