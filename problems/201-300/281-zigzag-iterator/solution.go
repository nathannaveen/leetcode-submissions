type ZigzagIterator struct {
    pos1, pos2 int
    v1, v2 []int
    v1Turn bool
}

func Constructor(v1, v2 []int) *ZigzagIterator {
    return &ZigzagIterator{0, 0, v1, v2, false}
}

func (this *ZigzagIterator) next() int {
    if this.pos1 < len(this.v1) && this.pos2 < len(this.v2) {
        
        this.v1Turn = !this.v1Turn
        if this.v1Turn {
            this.pos1++
            return this.v1[this.pos1 - 1]
        } else {
            this.pos2++
            return this.v2[this.pos2 - 1]
        }
    } else if this.pos1 < len(this.v1) {
        this.pos1++
        return this.v1[this.pos1 - 1]
    } else {
        this.pos2++
        return this.v2[this.pos2 - 1]
    }
}

func (this *ZigzagIterator) hasNext() bool {
	return this.pos1 < len(this.v1) || this.pos2 < len(this.v2)
}

/**
 * Your ZigzagIterator object will be instantiated and called as such:
 * obj := Constructor(param_1, param_2);
 * for obj.hasNext() {
 *	 ans = append(ans, obj.next())
 * }
 */