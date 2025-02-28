type key struct {
    i int // index in ingredients
    food string
}

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
    sups := make(map[string] bool)
    queue := []key{}
    res := []string{}
    
    for _, supply := range supplies {
        sups[supply] = true
    }
    
    for i, recipe := range recipes {
        queue = append(queue, key{ i, recipe })
    }
    
    infiniteLoop := false
    
    for len(queue) != 0 && !infiniteLoop {
        n := len(queue)
        infiniteLoop = true
        
        for l := 0; l < n; l++ {
            pop := queue[0]
            queue = queue[1:]
            canMake := true
            
            for _, ingredient := range ingredients[pop.i] {
                if !sups[ingredient] {
                    canMake = false
                    break
                }
            }
            
            if canMake {
                sups[pop.food] = true
                res = append(res, pop.food)
                infiniteLoop = false
            } else {
                queue = append(queue, pop)
            }
            
        }
    }
    
    return res
}