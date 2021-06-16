func minFlips(a int, b int, c int) int {
	res := 0
	for a > 0 || b > 0 || c > 0 {
		if c & 1 == 0 {
			if a & 1 == 1 {
				res++
			}
			if b & 1 == 1 { 
				res++ 
			}
		} else if a & 1 == 0 && b & 1 == 0 {
			res++
		}
		
		a, b, c = a >> 1, b >> 1, c >> 1
	}
	return res
}