type point struct {
	i, j int
}

func pacificAtlantic(heights [][]int) [][]int {
	n, m := len(heights), len(heights[0])

	pacificVisited := make(map[point] bool)
	atlanticVisited := make(map[point] bool)

	res := [][]int{}

	pacific, atlantic := findNextToWater(n, m)

	add := func(i, j, prevHeight int, pOrA string) {
		if pOrA == "pacific" {
			if i < 0 || j < 0 || i >= n || j >= m || pacificVisited[point{i, j}] {
				return
			}

			if heights[i][j] >= prevHeight {
				pacific = append(pacific, point{i, j})
			}
		} else { // Otherwise it is "atlantic"
			if i < 0 || j < 0 || i >= n || j >= m || atlanticVisited[point{i, j}] {
				return
			}

			if heights[i][j] >= prevHeight {
				atlantic = append(atlantic, point{i, j})
			}
		}
	}

	for len(pacific) != 0 {
        pop := pacific[len(pacific) - 1]
        pacific = pacific[:len(pacific) - 1]

		if pacificVisited[pop] {
			continue
		}

		pacificVisited[pop] = true
		i, j, h := pop.i, pop.j, heights[pop.i][pop.j]

		add(i+1, j, h, "pacific")
		add(i-1, j, h, "pacific")
		add(i, j+1, h, "pacific")
		add(i, j-1, h, "pacific")
	}

	for len(atlantic) != 0 {
        pop := atlantic[len(atlantic) - 1]
        atlantic = atlantic[:len(atlantic) - 1]

		if atlanticVisited[pop] {
			continue
		}

		if pacificVisited[pop] {
			res = append(res, []int{pop.i, pop.j})
		}

		atlanticVisited[pop] = true
		i, j, h := pop.i, pop.j, heights[pop.i][pop.j]

		add(i+1, j, h, "atlantic")
		add(i-1, j, h, "atlantic")
		add(i, j+1, h, "atlantic")
		add(i, j-1, h, "atlantic")
	}

	return res
}

func findNextToWater(n int, m int) ([]point, []point) {
	pacific := []point{}
	atlantic := []point{}

	for i := 0; i < n; i++ { // Find all values on the edge next to the ocean
		pacific = append(pacific, point{i, 0})
		atlantic = append(atlantic, point{i, m - 1})
	}
	for j := 1; j < m-1; j++ { // Find all values on the edge next to the ocean
		pacific = append(pacific, point{0, j})
		atlantic = append(atlantic, point{n - 1, j})
	}

	pacific = append(pacific, point{0, m - 1})
	atlantic = append(atlantic, point{n - 1, 0})
	return pacific, atlantic
}