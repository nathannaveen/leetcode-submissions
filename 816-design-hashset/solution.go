type MyHashSet struct {
    m map[int] int
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
    return MyHashSet{map[int]int{}}
}


func (this *MyHashSet) Add(key int)  {
    this.m[key] = 1
}


func (this *MyHashSet) Remove(key int)  {
    this.m[key] = 0
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
    return this.m[key] == 1
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */