type pos struct {
    i, j int
}

type TicTacToe struct {
    moves map[pos] int
    n int
}


func Constructor(n int) TicTacToe {
    return TicTacToe{map[pos] int {}, n}
}


func (this *TicTacToe) Move(row int, col int, player int) int {
    this.moves[pos{row, col}] = player

    for i := 0; i < this.n; i++ {
        canDo := true
        for j := 1; j < this.n; j++ {
            if this.moves[pos{i, j}] != this.moves[pos{i, j - 1}] {
                canDo = false
                break
            }
        }
        if canDo {
            return this.moves[pos{i, 0}]
        }
    }

    for i := 0; i < this.n; i++ {
        canDo := true
        for j := 1; j < this.n; j++ {
            if this.moves[pos{j, i}] != this.moves[pos{j - 1, i}] {
                canDo = false
                break
            }
        }
        if canDo {
            return this.moves[pos{0, i}]
        }
    }

    one := true
    two := true

    for i := 1; i < this.n; i++ {
        if this.moves[pos{i, i}] != this.moves[pos{i - 1, i - 1}] {
            one = false
        }
        if this.moves[pos{i, this.n - i - 1}] != this.moves[pos{i - 1, this.n - i}] {
            two = false
        }
    }

    if one {
        return this.moves[pos{0, 0}]
    }
    if two {
        return this.moves[pos{0, this.n - 1}]
    }

    return 0
}


/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */