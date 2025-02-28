type Node struct {
    m map[rune] *Node
    pos [][]int
}

type Trie struct {
    node *Node
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{node: &Node{make(map[rune] *Node), [][]int{}}}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string, start, end int) {
    root := this.node
    
    for _, i := range word {
        if root.m[i] == nil {
            root.m[i] = &Node{ m: make(map[rune] *Node), }
        }
        
        root = root.m[i]
    }
    root.pos = append(root.pos, []int{ start, end })
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) [][]int {
    root := this.node
    
    for _, i := range word {
        root = root.m[i]
        
        if root == nil {
            return [][]int{}
        }
    }
    return root.pos
}

func indexPairs(text string, words []string) [][]int {
    t := Constructor()
    res := [][]int{}
    
    for i := 1; i < len(text); i++ {
        s := text[0 : i]
        t.Insert(s, 0, i - 1)
        
        for j := i; j < len(text); j++ {
            s = s[1:]
            s += string(text[j])
            t.Insert(s, j - i + 1, j)
        }
    }
    
    for _, word := range words {
        for _, i := range t.Search(word) {
            res = append(res, i)
        }
    }
    
    sort.Slice(res, func(i, j int) bool {
        if res[i][0] == res[j][0] { return res[i][1] < res[j][1] }
        return res[i][0] < res[j][0]
    })
    
    return res
}





