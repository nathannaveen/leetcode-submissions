type NumArray struct {
    prefix []int
}


func Constructor(nums []int) NumArray {
    prefix := []int{0}

    for _, num := range nums {
        prefix = append(prefix, prefix[len(prefix) - 1] + num)
    }

    return NumArray{prefix}
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.prefix[right + 1] - this.prefix[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */