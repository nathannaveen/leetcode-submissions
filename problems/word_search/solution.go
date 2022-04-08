var b = [][]byte{}
var w = ""

func exist(board [][]byte, word string) bool {
    b = board
    w = word
    
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if board[i][j] == word[0] && helper(i, j, 0) {
                return true
            }
        }
    }
    
    return false
}

func helper(i, j, c int) bool {
    if i < 0 || j < 0 || i >= len(b) || j >= len(b[0]) || b[i][j] != w[c] {
        return false
    }
    
    if c == len(w) - 1 {
        return true
    }
    
    b[i][j] -= 33
    
    if helper(i + 1, j, c + 1) || helper(i - 1, j, c + 1) || helper(i, j + 1, c + 1) || helper(i, j - 1, c + 1) {
        return true
    }
    
    b[i][j] += 33
    
    return false
}