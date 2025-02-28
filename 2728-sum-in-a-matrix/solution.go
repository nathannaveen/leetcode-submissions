func matrixSum(nums [][]int) int {
   res := 0

    for _, row := range nums {
        sort.Slice(row, func(i, j int) bool { return row[i] > row[j] })
    }

    for len(nums[0]) > 0 {
        max := 0

        for i := 0; i < len(nums); i++ {
            if nums[i][0] > max {
                max = nums[i][0]
            }
            nums[i] = nums[i][1:]
        }

        res += max
    }

    return res 
}