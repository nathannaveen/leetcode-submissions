func orangesRotting(grid [][]int) int {
    rotten := [][]int{} // [][]int{ {i, j} }
    oranges := 0
    cnt := 0

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 2 {
                rotten = append(rotten, []int{i, j})
            }
            if grid[i][j] != 0 {
                oranges++
            }
        }
    }

    for len(rotten) > 0 {
        n := len(rotten)
        cnt++

        for c := 0; c < n; c++ {
            appendToQueue := func(i, j int) {
                if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] != 1 {
                    return
                }
                rotten = append(rotten, []int{i, j})
                grid[i][j] = 2
            }
            pop := rotten[0]
            rotten = rotten[1:]
            oranges--

            for _, d := range [][]int{ {0, 1}, {0, -1}, {1, 0}, {-1, 0} } {
                appendToQueue(pop[0] + d[0], pop[1] + d[1])
            }
        }
    }

    if oranges != 0 {
        return -1
    }

    return int(math.Max(float64(cnt - 1), float64(0)))
}