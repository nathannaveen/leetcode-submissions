func makeAntiPalindrome(s string) string {
    arr := strings.Split(s, "")

    freq := map[rune] int {}

    for _, l := range s {
        freq[l]++
        if freq[l] > len(s) / 2 {
            return "-1"
        }
    }

    sort.Strings(arr)

    prev := ""
    prevIndex := 0

    for i := len(arr) / 2; i < len(arr); i++ {
        x := i

        if arr[i] == prev {
            x = prevIndex
        }

        prev = arr[i]

        for arr[x] == arr[len(arr) - i - 1] {
            x++
        }

        arr[i], arr[x] = arr[x], arr[i]
        prevIndex = x
    }

    return strings.Join(arr, "")
}

/*
abccccdd

abcc
ddcc

*/
