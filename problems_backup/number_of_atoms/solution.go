type key struct {
    s string
    n int
}

func countOfAtoms(formula string) string {
    arr := []string{}
    stack := []key{}
    m := map[string] int {} // string : index in res2
    res2 := []key{}
    res := ""
    
    part := ""
    
    
    for _, l := range formula {
        if unicode.IsLower(l) || (unicode.IsDigit(l) && len(part) > 0 && unicode.IsDigit(rune(part[len(part) - 1]))) {
            part += string(l)
        } else {
            if part != "" {
                arr = append(arr, part)
            }
            part = string(l)
        }
        
    }
    
    arr = append(arr, part)
    
    for i := 0; i < len(arr); i++ {
        if arr[i] != ")" {
            if arr[i] == "(" {
                stack = append(stack, key{ arr[i], 1 })
            } else {
                if i == len(arr) - 1 {
                    stack = append(stack, key{ arr[i], 1 })
                } else {
                    n, err := strconv.Atoi(arr[i + 1])
                    if err == nil {
                        stack = append(stack, key{ arr[i], n })
                        i++
                    } else {
                        stack = append(stack, key{ arr[i], 1 })
                    }
                }
            }
        } else {
            a := []key{}

            n := 1

            newI := i

            if i != len(arr) - 1 {
                n2, err := strconv.Atoi(arr[i + 1])

                if err != nil {
                    n = 1
                } else {
                    n = n2
                    newI = i + 1
                }
            }

            for len(stack) > 0 {
                pop := stack[len(stack) - 1]
                stack = stack[:len(stack) - 1]

                if pop.s == "(" {
                    break
                }
                a = append(a, key{ pop.s, pop.n * n })
            }

            stack = append(stack, a...)
            i = newI
        }
    }
    
    for _, s := range stack {
        if _, ok := m[s.s]; !ok {
            res2 = append(res2, key{ s.s, 0 })
            m[s.s] = len(res2) - 1
        }
        res2[m[s.s]].n += s.n
    }
    
    sort.Slice(res2, func(i, j int) bool {
        return res2[i].s < res2[j].s
    })
    
    for _, r := range res2 {
        res += r.s
        if r.n != 1 {
            res += strconv.Itoa(r.n)
        }
    }
    
    return res
}