type FirstUnique struct {
    m map[int] int
    arr []int
}


func Constructor(nums []int) FirstUnique {
    m := make(map[int] int)
    arr := []int{}
    for i := 0; i < len(nums); i++ {
        m[nums[i]]++
        if m[nums[i]] > 1 {
            for j := 0; j < len(arr); j++ {
                if arr[j] == nums[i] {
                    arr = append(arr[:j], arr[j+1:]...)
                }
            }
        } else {
            arr = append(arr, nums[i])
        }
    }
    return FirstUnique{m, arr}
}


func (this *FirstUnique) ShowFirstUnique() int {
    if len(this.arr) == 0 { return - 1 }
    return this.arr[0]
}


func (this *FirstUnique) Add(value int)  {
    this.m[value]++
    if this.m[value] > 1 {
        for i := 0; i < len(this.arr); i++ {
            if this.arr[i] == value {
                this.arr = append(this.arr[:i], this.arr[i+1:]...)
            }
        }
    } else {
        this.arr = append(this.arr, value)
    }
}


/**
 * Your FirstUnique object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.ShowFirstUnique();
 * obj.Add(value);
 */