func tictactoe(moves [][]int) string {
	rows := make([]int, 3)
	cols := make([]int, 3)
	diag1 := 0
	diag2 := 0

	for i, move := range moves {
		if i % 2 == 0 {
			rows[move[1]]++
			cols[move[0]]++
			if move[0] == move[1] { diag1++ }
			if move[0] + move[1] == 2 { diag2++ }
		} else {
			rows[move[1]]--
			cols[move[0]]--
			if move[0] == move[1] { diag1-- }
			if move[0] + move[1] == 2 { diag2-- }
		}
	}
	for _, row := range rows {
		if row == 3 { return "A" }
		if row == -3 { return "B" }
	}
	for _, col := range cols {
		s, done := checking(col)
		if done { return s }
	}
	s, done := checking(diag1)
	if done { return s }
	s2, done2 := checking(diag2)
	if done2 { return s2 }

	if len(moves) == 9 {
		return "Draw"
	}
	return "Pending"
}

func checking(n int) (string, bool) {
	if n == 3 { return "A", true }
	if n == -3 { return "B", true }
	return "", false
}
