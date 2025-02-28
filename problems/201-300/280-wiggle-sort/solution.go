func wiggleSort(nums []int)  {
    sort.Ints(nums)
    for i := 1; i < len(nums) - 1; i += 2 {
        nums[i], nums[i + 1] = nums[i + 1], nums[i]
    }
}