func alphabetBoardPath(target string) string {
	board := []string{"abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"}
	res := ""
	row, col := 0, 0

	for i := 0; i < len(target); i++ {
		for newCol := 0; newCol < 6; newCol++ {
			newRow := containsPosition(board[newCol], int32(target[i]))
			if newRow != -1 { // this says that the row contains the letter
				amountOfY := abs(newCol - col)
				
				if newCol > col {
					for k := col; k < newCol; k++ {
						if k > len(board)-3 {
							i--
							break
						}
						res += "D"
						amountOfY--
					}
				} else if newCol < col {
					for k := 0; k < col-newCol; k++ {
						res += "U"
						amountOfY--
					}
				}
				
				if newRow > row {
					for k := 0; k < newRow-row; k++ {
						res += "R"
					}
				} else if newRow < row {
					for k := 0; k < row-newRow; k++ {
						res += "L"
					}
				}

				if amountOfY != 0 {
					res += "D"
				}
				if amountOfY == 0 {
					res += "!"
				}
				row, col = newRow, newCol
				break
			}
		}
	}

	return res
}

func containsPosition(theRow string, letter int32) int { 
	/*
	This checks whether theRow contains the letter 
	*/
	
	for i, i2 := range theRow {
		if i2 == letter {
			return i
		}
	}
	return -1
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}