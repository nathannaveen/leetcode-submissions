func numRookCaptures(board [][]byte) int {
	row, col := 0, 0

	for i, bytes := range board {
		gotRook := false
		for i2, b := range bytes {
			if b == 'R' {
				row = i
				col = i2
				gotRook = true
				break
			}
		}
		if gotRook {
			break
		}
	}

	res := 0

	for i := row + 1; i < len(board); i++ {
		if board[i][col] == 'p' {
			res++
			break
		}
		if board[i][col] == 'B' {
			break
		}
	}
	for i := row - 1; i >= 0; i-- {
		if board[i][col] == 'p' {
			res++
			break
		}
		if board[i][col] == 'B' {
			break
		}
	}
	for i := col + 1; i < len(board[0]); i++ {
		if board[row][i] == 'p' {
			res++
			break
		}
		if board[row][i] == 'B' {
			break
		}
	}
	for i := col - 1; i >= 0; i-- {
		if board[row][i] == 'p' {
			res++
			break
		}
		if board[row][i] == 'B' {
			break
		}
	}
	return res
}