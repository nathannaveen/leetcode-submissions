type FoodRatings struct {
    highest map[string] *KeyHeap
    actualRating map[string] int
    foodToCuisine map[string] string
}


func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
    highest := map[string] *KeyHeap {}
    actualRating := map[string] int {}
    foodToCuisine := map[string] string {}

    for i := 0; i < len(foods); i++ {
        if _, ok := highest[cuisines[i]]; !ok {
            highest[cuisines[i]] = &KeyHeap{}
            heap.Init(highest[cuisines[i]])
        }
        heap.Push(highest[cuisines[i]], key{foods[i], ratings[i]})
        actualRating[foods[i]] = ratings[i]
        foodToCuisine[foods[i]] = cuisines[i]
    }

    return FoodRatings{highest, actualRating, foodToCuisine}
}


func (this *FoodRatings) ChangeRating(food string, newRating int)  {
    heap.Push(this.highest[this.foodToCuisine[food]], key{food, newRating})
    this.actualRating[food] = newRating
}


func (this *FoodRatings) HighestRated(cuisine string) string {
    for (*this.highest[cuisine])[0].n != this.actualRating[(*this.highest[cuisine])[0].s] { 
        heap.Pop(this.highest[cuisine])
    }

    return (*this.highest[cuisine])[0].s
}

type key struct {
    s string
    n int
}

type KeyHeap []key

func (h KeyHeap) Len() int           { return len(h) }
func (h KeyHeap) Less(i, j int) bool { 
    if h[i].n == h[j].n {
        return h[i].s < h[j].s
    }
    return h[i].n > h[j].n 
}
func (h KeyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *KeyHeap) Push(x interface{}) {
	*h = append(*h, x.(key))
}

func (h *KeyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */