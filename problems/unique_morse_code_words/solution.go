func uniqueMorseRepresentations(words []string) int {
	letters := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",
		".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}
	m := make(map[string] int)

	for _, word := range words {
		s := ""
		for _, i := range word {
			s += letters[i - 'a']
		}
		m[s]++
	}
	return len(m)
}