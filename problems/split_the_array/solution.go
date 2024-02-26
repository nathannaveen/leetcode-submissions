func isPossibleToSplit(nums []int) bool {
    freq := map[int] int {}

    for _, n := range nums {
        freq[n]++

        if freq[n] > 2 {
            return false
        }
    }
    return true
}