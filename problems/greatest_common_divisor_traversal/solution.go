func canTraverseAllPairs(nums []int) bool {
    m := map[int] []int {} // prime factor : []int{num0, num1}
    
    for i := range nums {
        x := PrimeFactors(nums[i])
        
        for k := range x {
            m[k] = append(m[k], i)
        }
    }
    
    parents := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		parents[i] = i
	}
	u := &unionFind{
		count:   len(nums),
		parents: parents,
	}
    
    // fmt.Println(m)
    
    for _, v := range m {
        for i := 1; i < len(v); i++ {
            // fmt.Println(v[i], v[i - 1])
            u.Union(v[i], v[i - 1])
        }
    }
    
    // fmt.Println(u)
    
    if u.count == 1 {
        return true
    }
    return false
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

func PrimeFactors(n int) map[int] bool {
    pfs := map[int] bool {}
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs[2] = true
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
            pfs[i] = true
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
        pfs[n] = true
	}

	return pfs
}