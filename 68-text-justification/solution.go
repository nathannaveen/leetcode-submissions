func addSpaces(start, end int) (res string) {
    for i := start; i < end; i++ {
        res += " "
    }
    return res
}

func fullJustify(words []string, maxWidth int) []string {
    start := 0
    res := []string{}
    lineLength := 0
    
    for i := 0; i < len(words); i++ {
        if lineLength + len(words[i]) <= maxWidth {
            lineLength += len(words[i]) + 1
        } else {
            numberOfWords := i - start
            line := ""
            
            if numberOfWords == 1 {
                line = words[i - 1] + addSpaces(len(words[i - 1]), maxWidth)
            } else {
                lineLength -= 1
                numBigSpaces := numberOfWords - 1
                numberOfSpaces := (maxWidth - lineLength) / numBigSpaces
                additionalSpaces := (maxWidth - lineLength) % numBigSpaces
                
                j := start
                
                for numBigSpaces > 0 {
                    line += words[j] + " "
                    
                    if additionalSpaces > 0 {
                        additionalSpaces--
                        line += " "
                    }
                    
                    line += addSpaces(0, numberOfSpaces)
                    
                    j++
                    numBigSpaces--
                }
                line += words[i - 1]
            }
            
            res = append(res, line)
            lineLength = 0
            start = i
            i--
        }
    }
    
    // The last line
    
    line := ""
    
    for i := start; i < len(words); i++ {
        line += words[i] + " "
    }
    
    if len(line) == maxWidth + 1 {
        line = line[:len(line) - 1]
    }
    
    line += addSpaces(len(line), maxWidth)
    
    res = append(res, line)
    
    // return statment
    
    return res
}
