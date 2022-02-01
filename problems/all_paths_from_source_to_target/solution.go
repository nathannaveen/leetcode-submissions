var res = [][]int{}
var connections = make(map[int] []int)
var end = 0

func allPathsSourceTarget(graph [][]int) [][]int {
    res = [][]int{}
    end = len(graph) - 1
    connections = make(map[int] []int)
    
    for i, nodes := range graph {
        connections[i] = append(connections[i], nodes...)
    }
    
    helper(0, []int{})
    
    return res
}

func helper(node int, path []int) {
    path = append(path, node)
    
    if node == end {
        res = append(res, path)
        return
    }
    
    for _, n := range connections[node] {
        newPath := make([]int, len(path))
        copy(newPath, path)
        
        helper(n, newPath)
    }
}