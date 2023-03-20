func maximizeGreatness(nums []int) int {
    sort.Ints(nums)
    i := 0 // regular index
    j := 0 // above index
    res := 0

    for i < len(nums) && j < len(nums) {
        if nums[i] < nums[j] {
            i++
            res++
        }
        j++
    }

    return res
}