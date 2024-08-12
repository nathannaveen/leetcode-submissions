func minDays(grid [][]int) int {
    numberOfOnes := 0
    zeroCount := 0

    parents := make([]int, len(grid) * len(grid[0]))
	
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            parents[i * len(grid[0]) + j] = i * len(grid[0]) + j

            if grid[i][j] == 0 {
                zeroCount++
            } else {
                numberOfOnes++
            }
        }
    }

    newP := make([]int, len(grid) * len(grid[0]))
    copy(newP, parents)

    if cnt := helper(grid, newP); cnt - zeroCount != 1 {
        return 0
    }

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 0 {
                continue
            }

            newParents := make([]int, len(grid) * len(grid[0]))
            copy(newParents, parents)

            grid[i][j] = 0
            if count := helper(grid, newParents); count - zeroCount - 1 != 1 {
                return 1
            }
            grid[i][j] = 1
        }
    }
    
    return 2
}

func helper(grid [][]int, parents []int) int {
    uf := &unionFind{
		count:   len(parents),
		parents: parents,
	}

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 0 {
                continue
            }
            
            for _, d := range [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
                I, J := i + d[0], j + d[1]

                if inBounds(I, J, len(grid), len(grid[0])) && grid[I][J] == 1 {
                    uf.Union(i * len(grid[0]) + j, I * len(grid[0]) + J)
                }
            }
        }
    }

    return uf.count
}

type unionFind struct {
	count   int
	parents []int
}

func (this *unionFind) Find(x int) int {
	if x == this.parents[x] {
		return x
	}
	this.parents[x] = this.Find(this.parents[x])
	return this.parents[x]
}

func (this *unionFind) Union(x, y int) {
	rx, ry := this.Find(x), this.Find(y)
	if rx != ry {
		this.parents[rx] = ry
		this.count--
	}
}


func inBounds(i, j, I, J int) bool {
    return i >= 0 && j >= 0 && i < I && j < J
}

