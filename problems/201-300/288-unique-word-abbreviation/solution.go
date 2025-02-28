type ValidWordAbbr struct {
    m map[string] []string
}


func Constructor(dictionary []string) ValidWordAbbr {
    m := map[string] []string {}

    for _, d := range dictionary {
        abbreviation := wordToAbbr(d)
        m[abbreviation] = append(m[abbreviation], d)
    }

    return ValidWordAbbr{m}
}


func (this *ValidWordAbbr) IsUnique(word string) bool {
    val, ok := this.m[wordToAbbr(word)]
    found := false
    
    for _, v := range val {
        if v != word {
            return false
        } else {
            found = true
        }
    }
    return !ok || found
}


func wordToAbbr(word string) string {
    if len(word) < 3 {
        return word
    }
    return string(word[0]) + strconv.Itoa(len(word) - 2) + string(word[len(word) - 1])
}


/**
 * Your ValidWordAbbr object will be instantiated and called as such:
 * obj := Constructor(dictionary);
 * param_1 := obj.IsUnique(word);
 */