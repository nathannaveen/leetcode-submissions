func checkZeroOnes(s string) bool {
	oneCounter, zeroCounter := 0, 0
	resOne, resZero := 0, 0

	for _, i := range s {
		if i == '1' {
			oneCounter, zeroCounter = oneCounter + 1, 0
			resOne = max(oneCounter, resOne)
		} else {
			zeroCounter, oneCounter = zeroCounter + 1, 0
			resZero = max(zeroCounter, resZero)
		}
	}

	return resOne > resZero
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
