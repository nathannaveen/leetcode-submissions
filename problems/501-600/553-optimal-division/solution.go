func optimalDivision(nums []int) string {
    res := ""
    for i := 1; i < len(nums); i++ {
        res += strconv.Itoa(nums[i]) + "/" 
    }

    if len(res) > 2 {
        return strconv.Itoa(nums[0]) + "/(" + res[:len(res) - 1] + ")"
    } else if len(res) == 2 {
        return strconv.Itoa(nums[0]) + "/" + res[:len(res) - 1]
    } else {
        return strconv.Itoa(nums[0])
    }
}