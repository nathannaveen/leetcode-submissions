type key struct {
    i int
    j int
}

var connectedToEdge = make(map[key] bool)

func solve(board [][]byte)  {
    connectedToEdge = make(map[key] bool)
    
    for i := 0; i < len(board); i++ {
        if board[i][0] == 'O' { helper(i, 0, board) }
        if board[i][len(board[0]) - 1] == 'O' { helper(i, len(board[0]) - 1, board) }
    }
    
    for j := 0; j < len(board[0]); j++ {
        if board[0][j] == 'O' { helper(0, j, board) }
        if board[len(board) - 1][j] == 'O' { helper(len(board) - 1, j, board) }
    }
    
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if !connectedToEdge[key{ i, j }] {
                board[i][j] = 'X'
            }
        }
    }
}

func helper(i, j int, board [][]byte) {
    if i >= 0 && j >= 0 && i < len(board) && j < len(board[0]) && !connectedToEdge[key{ i, j }] && board[i][j] == 'O' {
        connectedToEdge[key{ i, j }] = true
        helper(i + 1, j, board)
        helper(i - 1, j, board)
        helper(i, j + 1, board)
        helper(i, j - 1, board)
    }
}
