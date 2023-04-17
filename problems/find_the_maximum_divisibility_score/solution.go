func maxDivScore(nums []int, divisors []int) int {
    score := 0
    res := divisors[0]

    for i := 0; i < len(divisors); i++ {
        n := 0

        for j := 0; j < len(nums); j++ {
            if nums[j] % divisors[i] == 0 {
                n++
            }
        }

        if n > score {
            score = n
            res = divisors[i]
        } else if n == score && divisors[i] < res {
            res = divisors[i]
        }
    }

    return res
}