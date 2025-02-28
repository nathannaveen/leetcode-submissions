type SparseVector struct {
    m map[int] int // index, value
}

func Constructor(nums []int) SparseVector {
    m := make(map[int] int)
    for i := range nums {
        if nums[i] != 0 {
            m[i] = nums[i]
        }
    }
    return SparseVector{ m }
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    res := 0
    
    for a, b := range vec.m {
        if val, ok := this.m[a]; ok {
            res += (b * val)
        }
    }
    
    return res
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */