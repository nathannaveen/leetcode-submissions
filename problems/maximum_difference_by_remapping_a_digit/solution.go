func minMaxDifference(num int) int {
    digits, remapBig := []int{}, 0

    for ; num != 0; num /= 10 {
        digit := num % 10
        digits = append(digits, digit)
        if digit != 9 { remapBig = digit }
    }

    remapSmall := digits[len(digits) - 1]

    min, max := 0, 0

    for i := len(digits) - 1; i >= 0; i-- {
        min, max = min * 10, max * 10

        if digits[i] == remapSmall {
            min += 0
        } else {
            min += digits[i]
        }

        if digits[i] == remapBig {
            max += 9
        } else {
            max += digits[i]
        }
    }

    return max - min
}