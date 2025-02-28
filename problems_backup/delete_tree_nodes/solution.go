type key struct {
    index int
    val int
}

func deleteTreeNodes(nodes int, parent []int, value []int) int {
    children := map[int] []key {}

    for i := 0; i < nodes; i++ {
        children[parent[i]] = append(children[parent[i]], key{i, value[i]})
    }

    _, b, _ := dfs(children, 0, value[0])

    return nodes - b
}

func dfs(children map[int] []key, index, val int) (int, int, int) {
    sum := val // the sum of all values in this subtree
    removed := 0   // the number of nodes that we know we need to remove in this subtree
    nodesInSubtree := 1  // the number of nodes in the subtree

    for _, node := range children[index] {
        a, b, c := dfs(children, node.index, node.val)

        sum += a
        removed += b
        nodesInSubtree += c
    }

    if sum == 0 {
        return 0, nodesInSubtree, nodesInSubtree
    }

    return sum, removed, nodesInSubtree
}