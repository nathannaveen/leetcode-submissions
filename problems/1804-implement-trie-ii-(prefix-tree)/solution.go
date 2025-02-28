type Node struct {
    m map[rune] *Node
    isEndCounter int
    starts int
}

type Trie struct {
    node *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{node: &Node{make(map[rune] *Node), 0, 0}}
}


func (this *Trie) Insert(word string)  {
    root := this.node
    root.starts++
    
    for _, s := range word {
        if root.m[s] == nil {
            root.m[s] = &Node{ m: make(map[rune] *Node), }
        }
        
        root = root.m[s]
        root.starts++
    }
    root.isEndCounter++
}


func (this *Trie) CountWordsEqualTo(word string) int {
    root := this.node
    
    for _, s := range word {
        root = root.m[s]
        
        if root == nil || root.starts == 0 {
            return 0
        }
    }
    return root.isEndCounter
}


func (this *Trie) CountWordsStartingWith(prefix string) int {
    root := this.node
    
    for _, s := range prefix {
        root = root.m[s]
        
        if root == nil || root.starts == 0 {
            return 0
        }
    }
    return root.starts
}


func (this *Trie) Erase(word string)  {
    root := this.node
    root.starts--
    
    for _, s := range word {
        root = root.m[s]
        root.starts--
    }
    root.isEndCounter--
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.CountWordsEqualTo(word);
 * param_3 := obj.CountWordsStartingWith(prefix);
 * obj.Erase(word);
 */