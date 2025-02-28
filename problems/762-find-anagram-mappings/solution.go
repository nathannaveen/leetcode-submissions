func anagramMappings(A []int, B []int) []int {
	m := make(map[int] int)
	res := make([]int, len(A))

	for i, i2 := range B {
		m[i2] = i
	}

	for i := range A {
		res[i] = m[A[i]]
	}
    
	return res
}