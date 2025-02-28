type key struct {
    i, j int
}

type neighborSum struct {
    grid [][]int
    m map[int] key
}


func Constructor(grid [][]int) neighborSum {
    m := map[int] key {}

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            m[grid[i][j]] = key{i, j}
        }
    }

    return neighborSum{grid, m}
}


func (this *neighborSum) AdjacentSum(value int) int {
    sum := 0

    for _, d := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
        if inBounds(this.m[value].i + d[0], this.m[value].j + d[1], len(this.grid), len(this.grid)) {
            sum += this.grid[this.m[value].i + d[0]][this.m[value].j + d[1]]
        }
    }

    return sum
}


func (this *neighborSum) DiagonalSum(value int) int {
    sum := 0

    for _, d := range [][]int{{1, 1}, {-1, 1}, {1, -1}, {-1, -1}} {
        if inBounds(this.m[value].i + d[0], this.m[value].j + d[1], len(this.grid), len(this.grid)) {
            sum += this.grid[this.m[value].i + d[0]][this.m[value].j + d[1]]
        }
    }

    return sum
}

func inBounds(i, j, I, J int) bool {
    return i >= 0 && j >= 0 && i < I && j < J
}


/**
 * Your neighborSum object will be instantiated and called as such:
 * obj := Constructor(grid);
 * param_1 := obj.AdjacentSum(value);
 * param_2 := obj.DiagonalSum(value);
 */
