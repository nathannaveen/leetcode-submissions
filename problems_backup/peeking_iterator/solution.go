/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
    iter *Iterator
    i int
    arr []int
}

func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{iter, 0, []int{}}
}

func (this *PeekingIterator) hasNext() bool {
    return this.iter.hasNext() || this.i < len(this.arr)
}

func (this *PeekingIterator) next() int {
    if this.i < len(this.arr) {
        this.i++
        return this.arr[this.i - 1]
    } else {
        v := this.iter.next()
        this.arr = append(this.arr, v)
        this.i++
        return v
    }
}

func (this *PeekingIterator) peek() int {
    if this.i < len(this.arr) {
        return this.arr[this.i]
    } else {
        v := this.iter.next()
        this.arr = append(this.arr, v)
        return v
    }
}