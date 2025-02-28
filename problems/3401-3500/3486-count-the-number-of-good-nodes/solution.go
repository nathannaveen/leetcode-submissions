func countGoodNodes(edges [][]int) int {
    m := map[int] []int {}
    for _, edge := range edges {
        m[edge[0]] = append(m[edge[0]], edge[1])
        m[edge[1]] = append(m[edge[1]], edge[0])
    }

    _, res := helper(m, 0, 0)

    return res
}

func helper(m map[int] []int, node int, prev int) (int, int) {
    // return the number of nodes, and the number of valid subtrees

    cnt := -1
    totalCnt := 0
    isValid := true
    totalValid := 0

    for _, n := range m[node] {
        if n == prev {
            continue
        }

        numNodes, valid := helper(m, n, node)
        if cnt != -1 && cnt != numNodes {
            isValid = false
        }
        cnt = numNodes
        totalCnt += numNodes
        totalValid += valid
    }

    if isValid {
        totalValid++
    }
    return totalCnt + 1, totalValid
}
