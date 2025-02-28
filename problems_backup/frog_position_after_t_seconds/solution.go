type key struct {
    cur int
    probability float64
    visited map[int] bool
}

func frogPosition(n int, edges [][]int, t int, target int) float64 {
    m := map[int] []int {}

    for _, edge := range edges {
        m[edge[0]] = append(m[edge[0]], edge[1])
        m[edge[1]] = append(m[edge[1]], edge[0])
    }

    queue := []key{ {1, 1, map[int] bool { 1 : true }} }

    for len(queue) > 0 && t >= 0 {
        n := len(queue)

        for l := 0; l < n; l++ {
            pop := queue[0]
            queue = queue[1:]

            denominator := 0

            for _, node := range m[pop.cur] {
                if !pop.visited[node] {
                    denominator++
                }
            }

            if pop.cur == target && (t == 0 || denominator == 0) {
                return pop.probability
            }

            for _, node := range m[pop.cur] {
                if pop.visited[node] {
                    continue
                }
                v := newVisited(pop.visited, node)

                queue = append(queue, key{ 
                    node, 
                    pop.probability * (float64(1) / float64(denominator)), 
                    v,
                })
            }
        }

        t--
    }

    return float64(0)
}

func newVisited(visited map[int] bool, node int) map[int] bool {
    newVisited := map[int] bool { node : true }
    for k, v := range visited {
        newVisited[k] = v
    }
    return newVisited
}