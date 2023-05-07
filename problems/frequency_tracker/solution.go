type FrequencyTracker struct {
    f map[int] int
    m map[int] int
}


func Constructor() FrequencyTracker {
    return FrequencyTracker{map[int] int {}, map[int] int {}}
}


func (this *FrequencyTracker) Add(number int)  {
    this.f[number]++
    this.m[this.f[number] - 1]--
    this.m[this.f[number]]++
}


func (this *FrequencyTracker) DeleteOne(number int)  {
    if this.f[number] > 0 {
        this.f[number]--
        this.m[this.f[number] + 1]--
        this.m[this.f[number]]++
    }
}


func (this *FrequencyTracker) HasFrequency(frequency int) bool {
    if this.m[frequency] > 0 {
        return true
    }
    return false
}


/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */