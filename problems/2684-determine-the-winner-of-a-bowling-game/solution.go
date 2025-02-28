func isWinner(player1 []int, player2 []int) int {
    score1 := player1[0]
    
    if len(player1) > 1 {
        if player1[0] == 10 {
            score1 += player1[1] * 2
        } else {
            score1 += player1[1]
        }
    }
    
    for i := 2; i < len(player1); i++ {
        if player1[i - 1] == 10 || player1[i - 2] == 10 {
            score1 += player1[i] * 2
        } else {
            score1 += player1[i]
        }
    }
    
    score2 := player2[0]
    
    if len(player1) > 1 {
        if player2[0] == 10 {
            score2 += player2[1] * 2
        } else {
            score2 += player2[1]
        }
    }
    
    for i := 2; i < len(player2); i++ {
        if player2[i - 1] == 10 || player2[i - 2] == 10 {
            score2 += player2[i] * 2
        } else {
            score2 += player2[i]
        }
    }
    
    if score1 == score2 {
        return 0
    } else if score1 > score2 {
        return 1
    }
    
    return 2
}