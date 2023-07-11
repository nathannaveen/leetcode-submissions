func similarPairs(words []string) int {
    m := map[[26]int]int{}

    for i := 0; i < len(words); i++{
        curr := words[i]
        inside :=[26]int{}
        for m:=0;m<len(curr);m++{
            inside[curr[m]-'a'] = 1
        }
        m[inside]++
    }
    count :=0

    for _,value := range m {
        x := value*(value-1)/2
        count += x
    }
    return count
}