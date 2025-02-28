type PhoneDirectory struct {
    m map[int] bool
}


func Constructor(maxNumbers int) PhoneDirectory {
    m := map[int] bool {}
    for i := 0; i < maxNumbers; i++ {
        m[i] = true
    }
    return PhoneDirectory{ m }
}


func (this *PhoneDirectory) Get() int {
    for key := range this.m {
        delete(this.m, key)
        return key
    }
    return -1
}


func (this *PhoneDirectory) Check(number int) bool {
    return this.m[number]
}


func (this *PhoneDirectory) Release(number int)  {
    this.m[number] = true
}


/**
 * Your PhoneDirectory object will be instantiated and called as such:
 * obj := Constructor(maxNumbers);
 * param_1 := obj.Get();
 * param_2 := obj.Check(number);
 * obj.Release(number);
 */