func maxVowels(s string, k int) int {
	max := float64(0)
	h := strings.Split(s, "")
	
	for i := 0; i < k; i++ {
		if h[i] == "a" || h[i] == "e" || h[i] == "i" || h[i] == "o" || h[i] == "u" {
			max++
		}
	}
	
	counter := max

	for i := k; i < len(s); i++ {
		if h[i - k] == "a" || h[i - k] == "e" || h[i - k] == "i" || h[i - k] == "o" || h[i - k] == "u" {
			counter--
		}
		if h[i] == "a" || h[i] == "e" || h[i] == "i" || h[i] == "o" || h[i] == "u" {
			counter++
		}
		max = math.Max(float64(max), float64(counter))
	}
	
	return int(max)
}