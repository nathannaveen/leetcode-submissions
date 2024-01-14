func maxFrequencyElements(nums []int) int {
    freq := map[int] int {}
    max := 0
    total := 0

    for _, n := range nums {
        freq[n]++

        if freq[n] > max {
            max = freq[n]
            total = freq[n]
        } else if freq[n] == max {
            total += freq[n]
        }
    }

    return total
}