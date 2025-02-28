type Node struct {
    m map[rune] *Node
    sum int
}

type Trie struct {
    node *Node
}

/** Initialize your data structure here. */
func Constructor2() Trie {
    return Trie{node: &Node{make(map[rune] *Node), 0}}
}


/** Inserts a word into the trie. */
func (this *Trie) InsertWordWithVal(word string, val int, m map[string] int)  {     
    root := this.node
    add := val
    if m[word] != 0 {
        add -= m[word]
    }
                                                                               
    root.sum += add

    for _, s := range word {
        if root.m[s] == nil {
            root.m[s] = &Node{ m: make(map[rune] *Node), }
        }
        root = root.m[s]
        root.sum += add
    }
    
    m[word] = val
}

func (this *Trie) StartsWith(prefix string) int {
    root := this.node
    
    for _, i := range prefix {
        root = root.m[i]
        
        if root == nil {
            return 0
        }
    }
    return root.sum
}

type MapSum struct {
    t Trie
    m map[string] int
}


/** Initialize your data structure here. */
func Constructor() MapSum {
    return MapSum{ Constructor2(), make(map[string] int) }
}


func (this *MapSum) Insert(key string, val int)  {
    this.t.InsertWordWithVal(key, val, this.m)
}


func (this *MapSum) Sum(prefix string) int {
    return this.t.StartsWith(prefix)
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */