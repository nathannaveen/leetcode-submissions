func countBattleships(board [][]byte) int {
    res := 0
    
    for i := range board {
        for j := range board[i] {
            if board[i][j] == 'X' && 
            (i == len(board) - 1 || board[i + 1][j] == '.') && 
            (j == len(board[0]) - 1 || board[i][j + 1] == '.') {
                res++
            }
        }
    }
    
    return res
}