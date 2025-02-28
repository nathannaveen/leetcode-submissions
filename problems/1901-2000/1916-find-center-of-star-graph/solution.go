func findCenter(edges [][]int) int {
	nodes := make([]int, len(edges) + 2)

	for _, edge := range edges {
		nodes[edge[0]]++
		nodes[edge[1]]++
		if nodes[edge[0]] == 2 {
			return edge[0]
		}
		if nodes[edge[1]] == 2 {
			return edge[1]
		}
	}
	return edges[0][0]
}