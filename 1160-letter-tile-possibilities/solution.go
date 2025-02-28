var done = map[string] bool {}

func numTilePossibilities(tiles string) int {
    done = map[string] bool {}
    return helper("", tiles) - 1
}

func helper(s, tiles string) int {
    if done[s] { return 0 }
    
    res := 1
    done[s] = true
    
    for i := 0; i < len(tiles); i++ {
        res += helper(s + string(tiles[i]), tiles[:i] + tiles[i + 1:])
    }
    
    return res
}