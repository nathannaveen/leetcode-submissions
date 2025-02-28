func squareIsWhite(coordinates string) bool {
	letters := make(map[string]string)
	theNumber, _ := strconv.Atoi(string(coordinates[1]))
	theLetter := string(coordinates[0])

	letters["a"] = "even"; letters["c"] = "even"; letters["e"] = "even"; letters["g"] = "even"
	letters["b"] = "odd"; letters["d"] = "odd" ; letters["f"] = "odd"; letters["h"] = "odd"

	if letters[theLetter] == "even" && theNumber % 2 == 0 {
		return true
	} else if letters[theLetter] == "even" {
		return false
	}

	if theNumber % 2 == 0 {
		return false
	}
	return true
}