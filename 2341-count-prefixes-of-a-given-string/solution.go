type Node struct {
    m map[rune] *Node
}

type Trie struct {
    node *Node
}

func (this *Trie) Insert(word string)  {
    root := this.node
    
    for _, s := range word {
        if root.m[s] == nil {
            root.m[s] = &Node{ m: make(map[rune] *Node), }
        }
        
        root = root.m[s]
    }
}

func (this *Trie) isPrefix(word string) int {
    root := this.node
    
    for _, s := range word {
        root = root.m[s]
        
        if root == nil {
            return 0
        }
    }
    
    return 1
}

func countPrefixes(words []string, s string) int {
    t := Trie{&Node{make(map[rune] *Node)}}
    
    t.Insert(s)
    
    res := 0
    
    for _, word := range words {
        res += t.isPrefix(word)
    }
    
    return res
}