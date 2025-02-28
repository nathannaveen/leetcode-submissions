func judgeSquareSum(c int) bool {
	for i := 0; i*i <= c; i++ {
		if isSquare(c - (i * i)) {
			return true
		}
	}
	return false
}

func isSquare(n int) bool {
	squareRooted := int(math.Sqrt(float64(n)))
	if squareRooted*squareRooted == n {
		return true
	}
	return false
}