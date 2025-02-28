type MyHashMap struct {
    keys []int
    values []int
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{[]int{}, []int{}}
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
    contains := false
    for i := range this.keys {
        if this.keys[i] == key {
            contains = true
            this.values[i] = value
        }
    }
    if !contains {
        this.keys = append(this.keys, key)
        this.values = append(this.values, value)
    }
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    
    for i := range this.keys {
        if this.keys[i] == key {
            return this.values[i]
        }
    }
    
    return - 1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    for i := range this.keys {
        if this.keys[i] == key {
            this.keys = append(this.keys[:i], this.keys[i+1:]...)
            this.values = append(this.values[:i], this.values[i+1:]...)
            break
        }
    }
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */