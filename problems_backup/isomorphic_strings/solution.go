func isIsomorphic(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	amap, bmap := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(a); i++ {
		g, aok := amap[a[i]]
		h, bok := bmap[b[i]]
		if aok != bok || g != h {
			return false
		}
		amap[a[i]] = i + 1
		bmap[b[i]] = i + 1
	}
	return true
}