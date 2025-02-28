func dayOfYear(date string) int {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	res, _ := strconv.Atoi(date[8:])
	if month > 2 && year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		res++
	}
	res += 30 * (month - 1)
	if month > 2 {
		res -= 2
	}
    
	if month < 9 {
		res += month / 2
	} else {
		res += 4
		res += (month - 7) / 2
	}
	return res
}