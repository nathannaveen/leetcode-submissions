type TwoSum struct {
    arr []int
}


/** Initialize your data structure here. */
func Constructor() TwoSum {
    return TwoSum{[]int{}}
}


/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int)  {
    this.arr = append(this.arr, number)
}


/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
    sort.Ints(this.arr)
    left, right := 0, len(this.arr) - 1
    
    for left < right {
        if this.arr[left] + this.arr[right] < value {
            left++
        } else if this.arr[left] + this.arr[right] > value {
            right--
        } else {
            return true
        }
    }
    return false
}


/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */