func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	res := []int{}
	ratings := []int{}

	for _, restaurant := range restaurants {
		if veganFriendly == 1 && restaurant[2] == 0 {
			continue
		}
		if restaurant[3] > maxPrice {
			continue
		}
		if restaurant[4] > maxDistance {
			continue
		}

		res = append(res, restaurant[0])
		ratings = append(ratings, restaurant[1])
	}

	for i := 1; i < len(ratings); i++ {
		if i >= 1 && ratings[i-1] < ratings[i] {
			ratings[i - 1], ratings[i] = ratings[i], ratings[i - 1]
			res[i - 1], res[i] = res[i], res[i - 1]
			i-=2
		}
		if i >= 1 && ratings[i-1] == ratings[i] && res[i - 1] < res[i] {
			ratings[i - 1], ratings[i] = ratings[i], ratings[i - 1]
			res[i - 1], res[i] = res[i], res[i - 1]
			i-=2
		}
	}
	return res
}