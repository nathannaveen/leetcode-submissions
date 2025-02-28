type Node struct {
    m map[rune] *Node
    cnt int
}

type Trie struct {
    node *Node
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    root := this.node
    
    for _, s := range word {
        if root.m[s] == nil {
            root.m[s] = &Node{ make(map[rune] *Node), 0 }
        }
        
        root.m[s].cnt++
        root = root.m[s]
    }
}

func (this *Trie) StartsAndPrefix(word string) int {
    root := this.node
    res := 0
    
    for _, s := range word {
        if root == nil {
            return 0
        }
        
        root = root.m[s]
        res += root.cnt
    }
    
    return res
}

func sumPrefixScores(words []string) []int {
    trie := &Trie{ &Node{ make(map[rune] *Node), 0 }}
    res := make([]int, len(words))
    
    for _, word := range words {
        trie.Insert(word)
    }
    
    for i := 0; i < len(words); i++ {        
        res[i] = trie.StartsAndPrefix(words[i])
    }
    
    return res
}


