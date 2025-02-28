func stoneGame(piles []int) bool {
	left, right := 0, len(piles)-1
	alex, lee := 0, 0
	whosPlaying := true

	for left < right {
		if whosPlaying && piles[left] > piles[right] {
			alex += piles[left]
			left++
		} else if whosPlaying && piles[left] <= piles[right] {
			alex += piles[right]
			right--
		} else if !whosPlaying && piles[left] > piles[right] {
			lee += piles[right]
			right--
		} else if !whosPlaying && piles[left] <= piles[right] {
			lee += piles[left]
		}

		whosPlaying = !whosPlaying
	}
	return alex > lee
}