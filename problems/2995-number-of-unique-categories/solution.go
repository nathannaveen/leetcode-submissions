/**
 * Definition for a category handler.
 * type CategoryHandler interface {
 *  HaveSameCategory(int, int) bool
 * }
 */
func numberOfCategories(n int, categoryHandler CategoryHandler) int {
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}
	u := &unionFind{
		count:   n,
		parents: parents,
	}

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            if categoryHandler.HaveSameCategory(i, j) {
                u.Union(i, j)
            }
        }
    }

    return u.count
}

type unionFind struct {
	count   int
	parents []int
}

func (this *unionFind) Find(x int) int {
	if x == this.parents[x] {
		return x
	}
	this.parents[x] = this.Find(this.parents[x])
	return this.parents[x]
}

func (this *unionFind) Union(x, y int) {
	rx, ry := this.Find(x), this.Find(y)
	if rx != ry {
		this.parents[rx] = ry
		this.count--
	}
}