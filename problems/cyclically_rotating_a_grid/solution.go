type key struct {
    val int
    i, j int
}

func rotateGrid(grid [][]int, k int) [][]int {
    n, m := len(grid), len(grid[0])
    
    for l := 0; l < int(math.Min(float64(n), float64(m))) / 2; l++ {
        arr := []key{ }
        
        for i := l; i < n - l; i++ {
            arr = append(arr, key{ grid[i][l], i, l }, key{ grid[i][m - l - 1], i, m - l - 1 })
        }
        for i := l + 1; i < m - l - 1; i++ {
            arr = append(arr, key{ grid[l][i], l, i }, key{ grid[n - l - 1][i], n - l - 1, i })
        }

        for _, a := range arr {
            i := a.i
            j := a.j
            newK := k % (2 * (n - l - l - 1) + 2 * (m - l - l - 1))
            
            
            for newK > 0 {
                if (i == l && j == l) || (j == l && i != n - 1 - l) {
                    if i + newK > n - 1 - l {
                        newK -= ((n - 1 - l) - i)
                        i = n - 1 - l
                    } else {
                        i += newK
                        newK = 0
                    }
                } else if (i == n - 1 - l && j == m - 1 - l) || (j == m - 1 - l && i != l) {
                    if i - newK < l {
                        newK -= (i - l)
                        i = l
                    } else {
                        i -= newK
                        newK = 0
                    }
                } else if (i == n - 1 - l && j == l) || (i == n - 1 - l && j != m - 1 - l) {
                    if j + newK > m - 1 - l {
                        newK -= ((m - 1 - l) - j)
                        j = m - 1 - l
                    } else {
                        j += newK
                        newK = 0
                    }
                } else if (i == l && j == m - 1 - l) || (i == l && j != l) {
                    if j - newK < l {
                        newK -= (j - l)
                        j = l
                    } else {
                        j -= newK
                        newK = 0
                    }
                }
            }
            grid[i][j] = a.val
        }
    }
    
    return grid
}