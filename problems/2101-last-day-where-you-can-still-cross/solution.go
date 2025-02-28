type pos struct {
    i, j int
}

func latestDayToCross(row int, col int, cells [][]int) int {
    parents := make(map[pos] pos)
    columns := make(map[pos] map[int] bool)
	
    for i := 0; i < len(cells); i++ {
        cur := pos{cells[i][0], cells[i][1]}
		parents[cur] = cur
        columns[cur] = map[int] bool {}
	}

	u := &unionFind{
		parents: parents,
        columns: columns,
	}

    m := map[pos] bool {}
    directions := []pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, 1}, {1, -1}, {-1, -1}} // directions for adjacents and diagonals

    for i, cell := range cells {
        cell1 := pos{cell[0], cell[1]}
        m[cell1] = true

        for _, d := range directions {
            cell2 := pos{cell[0] + d.i, cell[1] + d.j}

            if m[cell2] {
                l := u.Union(cell1, cell2)

                if l == col {
                    return i
                }
            }
        }
    }

    return 0
}

type unionFind struct {
	parents map[pos] pos
    columns map[pos] map[int] bool
}

func (this *unionFind) Find(x pos) pos {
	if x == this.parents[x] {
		return x
	}
	this.parents[x] = this.Find(this.parents[x])
	return this.parents[x]
}

func (this *unionFind) Union(x, y pos) int {
	rx := this.Find(x)
    ry := this.Find(y)
	if rx != ry {
		this.parents[rx] = ry
        this.columns[ry][x.j] = true
        this.columns[ry][y.j] = true
        for k, v := range this.columns[rx] {
            this.columns[ry][k] = v
        }
	}

    return len(this.columns[ry])
}
