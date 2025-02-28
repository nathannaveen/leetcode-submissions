func numSmallerByFrequency(queries []string, words []string) []int {
    arr := []int{}
    res := []int{}

    for _, word := range words {
        arr = append(arr, f(word))
    }

    sort.Slice(arr, func(i, j int) bool {
        return arr[i] > arr[j]
    })

    for _, query := range queries {
        freq := f(query)
        cnt := 0

        for i := 0; i < len(arr); i++ {
            if arr[i] > freq {
                cnt++
            } else {
                break
            }
        }

        res = append(res, cnt)
    }

    return res
}

func f(s string) int {
    min := 'z'
    freq := 0
    
    for _, x := range s {
        if x < min {
            min = x
            freq = 1
        } else if x == min {
            freq++
        }
    }

    return freq
}