import "math"

func numSquarefulPerms(nums []int) int {
    return helper(nums, 0, 0, -1)
}

func helper(nums []int, used int, cnt int, prev int) int {
    if cnt == len(nums) {
        return 1
    }

    res := 0

    usedThisIndex := map[int] bool {}

    for i, num := range nums {
        if used & (1 << i) == 0 && !usedThisIndex[num] && (prev == -1 || isSquare(prev + num)) {
            res += helper(nums, used | (1 << i), cnt + 1, num)
            usedThisIndex[num] = true
        }
    }
    
    return res
}

func isSquare(a int) bool {
    x := math.Sqrt(float64(a))

    return x == float64(int(x))
}