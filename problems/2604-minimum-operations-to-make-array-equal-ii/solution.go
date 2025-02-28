func minOperations(nums1 []int, nums2 []int, k int) int64 {
    add, subtract := int64(0), int64(0 )
    // the number of elements that need to be added and subtracted

    for i := 0; i < len(nums1); i++ {
        diff := nums2[i] - nums1[i]
        
        if diff == 0 {
            continue
        }
        if k == 0 || diff % k != 0 {
            return -1
        }
        
        if diff < 0 {
            subtract += abs(diff / k)
        } else { // diff > 0
            add += abs(diff / k)
        }
    }

    if add != subtract {
        return -1
    }

    return add
}

func abs(a int) int64 {
    if a < 0 {
        return int64(-a)
    }
    return int64(a)
}