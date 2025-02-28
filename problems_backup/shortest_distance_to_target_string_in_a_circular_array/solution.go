func closetTarget(words []string, target string, startIndex int) int {
    indexes := []int{}

    for i, word := range words {
        if word == target {
            indexes = append(indexes, i)
        }
    }

    if len(indexes) == 0 {
        return -1
    }

    dist := 101

    for _, index := range indexes {
        a := abs(startIndex - index)
        b := abs(len(words) - abs(startIndex - index))
        
        if a < dist {
            dist = a
        }
        if b < dist {
            dist = b
        }
    }

    return dist
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}