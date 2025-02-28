func findIntersectionValues(nums1 []int, nums2 []int) []int {
    res := []int{0, 0}

    m1 := map[int] bool {}
    m2 := map[int] bool {}

    for _, n := range nums1 {
        m1[n] = true
    }

    for _, n := range nums2 {
        if m1[n] {
            res[1]++
        }
        m2[n] = true
    }

    for _, n := range nums1 {
        if m2[n] {
            res[0]++
        }
    }

    return res
}