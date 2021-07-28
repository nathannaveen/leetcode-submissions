type MagicDictionary struct {
    m map[int] []string // len(the word), the word
}


/** Initialize your data structure here. */
func Constructor() MagicDictionary {
    return MagicDictionary{ make(map[int] []string) }
}


func (this *MagicDictionary) BuildDict(dictionary []string)  {
    for _, s := range dictionary {
        this.m[len(s)] = append(this.m[len(s)], s)
    }
}


func (this *MagicDictionary) Search(searchWord string) bool {
    for _, s := range this.m[len(searchWord)] {
        counter := 0
        
        for i := 0; i < len(s); i++ {
            if s[i] != searchWord[i] {
                counter++
            }
        }
        if counter == 1 {
            return true
        }
    }
    
    return false
}


/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */