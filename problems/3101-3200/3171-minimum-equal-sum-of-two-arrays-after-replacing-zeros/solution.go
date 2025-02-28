func minSum(nums1 []int, nums2 []int) int64 {
    cnt1, zeroCnt1 := int64(0), 0
    cnt2, zeroCnt2 := int64(0), 0
    
    for _, n := range nums1 {
        if n == 0 {
            zeroCnt1++
        } else {
            cnt1 += int64(n)
        }
    }
    
    for _, n := range nums2 {
        if n == 0 {
            zeroCnt2++
        } else {
            cnt2 += int64(n)
        }
    }
    
    cnt1 += int64(zeroCnt1)
    cnt2 += int64(zeroCnt2)
    
    if cnt1 < cnt2 {
        if zeroCnt1 > 0 {
            return cnt2
        }
        return -1
    } else if cnt2 < cnt1 {
        if zeroCnt2 > 0 {
            return cnt1
        }
        return -1
    }
    return cnt1
}