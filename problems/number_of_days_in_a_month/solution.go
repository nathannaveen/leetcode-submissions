func numberOfDays(Y int, M int) int {
	if (M <= 7 && M % 2 == 1) || (M >= 8 && M % 2 == 0) { return 31 }
	if M != 2 { return 30 }
	if Y % 4 == 0 && (Y % 100 != 0 || Y % 400 == 0) { return 29 }
	return 28
}