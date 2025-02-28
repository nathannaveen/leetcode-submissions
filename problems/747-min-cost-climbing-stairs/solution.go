func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)
	h := []int{}
	h = append(h, cost[0], cost[1])
	for i := 2; i < len(cost); i++ {
		h = append(h, int(math.Min(float64(h[i-1]), float64(h[i-2])))+cost[i])
	}
	return h[len(cost)-1]
}