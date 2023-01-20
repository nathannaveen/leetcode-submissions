func oddString(words []string) string {
    for i := 0; i < len(words[0]); i++ {
        d := words[0][i] - words[0][0] // find the value when the first letter is reduced to 'a'
        d2 := words[1][i] - words[1][0]

        for j := 2; j < len(words); j++ {
            d3 := words[j][i] - words[j][0]
            if d == d2 && d3 != d {
                return words[j]
            }
            if d != d2 {
                if d3 != d {
                    return words[0]
                }
                if d3 != d2 {
                    return words[1]
                }
            }
        }
    }

   return words[0]
}