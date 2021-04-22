func getWinner(arr []int, k int) int {
	counter, winner := 0, arr[0]

	for i := 1; i < len(arr); i++ {
		if winner < arr[i] {
			winner = arr[i]
			counter = 0
		}
		counter++
		if counter == k {
			return winner
		}
	}
	return winner
}