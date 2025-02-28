func numberOfAlternatingGroups(colors []int, k int) int {
    full := false
    res := 0

    for i := 0; i < len(colors); i++ {
        if !full {
            if helper(colors, k, i) {
                full = true
                res++
            }
        } else {
            if colors[(i + k - 1) % len(colors)] != colors[(i + k - 2) % len(colors)] {
                res++
            } else {
                full = false
            }
        }
    }

    return res
}

func helper(colors []int, k int, start int) bool {
    prev := -1
    for i := start; i < start + k; i++ {
        if colors[i % len(colors)] == prev {
            return false
        }
        prev = colors[i % len(colors)]
    }

    return true
}
