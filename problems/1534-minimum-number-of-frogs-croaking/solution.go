func minNumberOfFrogs(croakOfFrogs string) int {
    croak := "croak"
    arr := []int{}
    max := 0

    for _, letter := range croakOfFrogs {
        if letter == 'c' {
            arr = append(arr, 1)
        } else {
            added := false
            for j := 0; j < len(arr); j++ {
                if letter == rune(croak[arr[j]]) {
                    arr[j]++
                    if arr[j] == 5 {
                        if len(arr) > max {
                            max = len(arr)
                        }
                        arr = append(arr[:j], arr[j + 1:]...)
                    }
                    added = true
                    break
                }
            }
            if !added {
                return -1
            }
        }
    }

    if len(arr) != 0 {
        return -1
    }

    return max
}