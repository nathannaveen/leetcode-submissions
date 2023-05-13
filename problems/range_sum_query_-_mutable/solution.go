type NumArray struct {
    lineSweep []int
}


func Constructor(nums []int) NumArray {
    lineSweep := make([]int, len(nums))
    copy(lineSweep, nums)

    return NumArray{lineSweep}
}


func (this *NumArray) Update(index int, val int)  {
    this.lineSweep[index] = val
}


func (this *NumArray) SumRange(left int, right int) int {
    lSum := 0

    for i := 0; i < left; i++ {
        lSum += this.lineSweep[i]
    }

    rSum := lSum

    for i := left; i <= right; i++ {
        rSum += this.lineSweep[i]
    }

    // subtract the left sum from the right sum to get the sum of 
    // values between left and right
    return rSum - lSum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */