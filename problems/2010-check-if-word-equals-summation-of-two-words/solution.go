func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	sum, sum2, res := "", "", ""

	for _, letter := range firstWord { sum += string(letter - '1') }
	for _, letter := range secondWord { sum2 += string(letter - '1') }
	for _, letter := range targetWord { res += string(letter - '1') }
    
	temp, _ := strconv.Atoi(sum)
	temp2, _ := strconv.Atoi(sum2)
	tempRes, _ := strconv.Atoi(res)

	return temp + temp2 == tempRes
}
