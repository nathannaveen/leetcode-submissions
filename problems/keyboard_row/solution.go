func findWords(words []string) []string {
	res := []string{}
	letters := make(map[string]int)
	letters["q"] = 0; letters["w"] = 1; letters["e"] = 2; letters["r"] = 3; letters["t"] = 4; letters["y"] = 5
	letters["u"] = 6; letters["i"] = 7; letters["o"] = 8; letters["p"] = 9; letters["a"] = 10; letters["s"] = 11
	letters["d"] = 12; letters["f"] = 13; letters["g"] = 14; letters["h"] = 15; letters["j"] = 16; letters["k"] = 17
	letters["l"] = 18; letters["z"] = 19; letters["x"] = 20; letters["c"] = 21; letters["v"] = 22; letters["b"] = 23
	letters["n"] = 24; letters["m"] = 25

	for _, word := range words {
		firstLetter := string(word[0])
		if word[0] <= 90 && word[0] >= 65 {
			firstLetter = string(word[0] + 32)
		}
		row := 0

		if letters[firstLetter] >= 10 && letters[firstLetter] <= 18 {
			row = 1
		} else if letters[firstLetter] >= 19 {
			row = 2
		}

		g := true
		for i := 1; i < len(word); i++ {
			letter := string(word[i])
			if word[i] <= 90 && word[i] >= 65 {
				letter = string(word[i] + 32)
			}
			
			if row == 0 && letters[letter] > 9 {
				g = false; break
			} else if row == 1 && (letters[letter] < 10 || letters[letter] > 18) {
				g = false; break
			} else if row == 2 && letters[letter] < 19 {
				g = false; break
			}
		}
		if g {
			res = append(res, word)
		}
	}

	return res
}