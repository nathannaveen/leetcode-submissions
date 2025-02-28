func separateDigits(nums []int) []int {
    res := []int{}

    for _, num := range nums {
        arr := []int{}
        for num > 0 {
            arr = append([]int{num % 10}, arr...)
            num /= 10
        }
        res = append(res, arr...)
    }

    return res
}