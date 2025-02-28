type Node struct {
    m map[rune] *Node
    isEnd bool
}

type Trie struct {
    node *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{node: &Node{make(map[rune] *Node), false}}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    root := this.node
    
    for _, s := range word {
        if root.m[s] == nil {
            root.m[s] = &Node{ m: make(map[rune] *Node), }
        }
        
        root = root.m[s]
    }
    root.isEnd = true
}

func (this *Trie) StartsAndPrefix(word string, isPrefix bool) bool {
    root := this.node
    
    for _, s := range word {
        root = root.m[s]
        
        if root == nil {
            return false
        }
    }
    if isPrefix {
        return true
    }
    return root.isEnd
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    return this.StartsAndPrefix(word, false)
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    return this.StartsAndPrefix(prefix, true)
}



/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */