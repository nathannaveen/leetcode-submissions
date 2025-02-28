func minSubArrayLen(target int, nums []int) int {
    start, end, sum := 0, 0, 0
    
    for end = 0; end < len(nums); end++ {
        sum += nums[end]

        if sum >= target {
            break
        }
    }

    if sum < target {
        return 0
    }

    min := end - start + 1

    for end < len(nums) {
        if sum >= target {
            if end - start < min {
                min = end - start + 1
            }
            sum -= nums[start]
            start++
        } else if end + 1 < len(nums) {
            end++
            sum += nums[end]
        } else {
            break
        }
    }

    return min
}