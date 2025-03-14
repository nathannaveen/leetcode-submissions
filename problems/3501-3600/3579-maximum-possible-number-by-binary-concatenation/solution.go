type key struct {
    a, b, c int
}

func maxGoodNumber(nums []int) int {
    m := map[int] int{}
    m[0] = len(strconv.FormatInt(int64(nums[0]), 2))
    m[1] = len(strconv.FormatInt(int64(nums[1]), 2))
    m[2] = len(strconv.FormatInt(int64(nums[2]), 2))
    res := 0
    
    for _, x := range []key{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}} {
        res = max(res, nums[x.a] << (m[x.b] + m[x.c]) + nums[x.b] << m[x.c] + nums[x.c])
    }
    
    return res
}
