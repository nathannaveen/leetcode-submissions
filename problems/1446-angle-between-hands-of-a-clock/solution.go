func angleClock(hour int, minutes int) float64 {
	hourHand := float64(minutes*30)/float64(60) + float64(hour%12*30)
	minHand := float64(6 * minutes)
	distance := math.Abs(hourHand - minHand)

	if distance > 180 {
		distance = 360 - distance
	}

	return distance
}
