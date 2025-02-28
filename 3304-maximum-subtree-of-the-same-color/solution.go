func maximumSubtreeSize(edges [][]int, colors []int) int {
    m := map[int] []int {}

    for _, edge := range edges {
        m[edge[0]] = append(m[edge[0]], edge[1])
    }

    _, _, res := helper(m, colors, 0)

    return res
}

func helper(m map[int] []int, colors []int, cur int) (int, int, int) {
    // return the color, the current size, and the maximum size
    color := colors[cur]
    size := 1
    maxSize := 0

    for _, node := range m[cur] {
        a, b, c := helper(m, colors, node)

        if a != color || a == -1 {
            color = -1
        }
        if c > maxSize {
            maxSize = c
        }
        
        size += b
    }

    if size > maxSize && color != -1 {
        maxSize = size
    }

    return color, size, maxSize
}