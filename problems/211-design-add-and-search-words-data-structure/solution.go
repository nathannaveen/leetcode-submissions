type WordDictionary struct {
    t *trie
}


func Constructor() WordDictionary {
    return WordDictionary{initTrie()}
}


func (this *WordDictionary) AddWord(word string)  {
    this.t.insert(word)
}


func (this *WordDictionary) Search(word string) bool {
    return this.t.find(word, 0)
}

type trieNode struct {
    childrens [26]*trieNode
    isWordEnd bool
}

type trie struct {
    root *trieNode
}

func initTrie() *trie {
    return &trie{
        root: &trieNode{},
    }
}

func (t *trie) insert(word string) {
    wordLength := len(word)
    current := t.root
    for i := 0; i < wordLength; i++ {
        index := word[i] - 'a'
        if current.childrens[index] == nil {
            current.childrens[index] = &trieNode{}
        }
        current = current.childrens[index]
    }
    current.isWordEnd = true
}

func (t *trie) find(word string, index int) bool {
    if t == nil || t.root == nil {
        return false
    }

    current := t.root

    if current.isWordEnd && index == len(word) {
        return true
    } else if index == len(word) {
        return false
    }
    letter := word[index]

    if letter == '.' {
        for i := 0; i < 26; i++ {
            newTrie := &trie{root: current.childrens[i]}
            
            if newTrie.find(word, index + 1) {
                return true
            }
        }
    } else {
        newTrie := &trie{root: current.childrens[letter - 'a']}
        
        if newTrie.find(word, index + 1) {
            return true
        }
    }

    return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */