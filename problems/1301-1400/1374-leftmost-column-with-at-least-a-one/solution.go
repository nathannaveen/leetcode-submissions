/**
 * // This is the BinaryMatrix's API interface.
 * // You should not implement it, or speculate about its implementation
 * type BinaryMatrix struct {
 *     Get func(int, int) int
 *     Dimensions func() []int
 * }
 */

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
    d := binaryMatrix.Dimensions()
    first := d[1] - 1
    found := false
    
    for i := 0; i < d[0]; i++ {
        for j := first; j >= 0; j-- {
            if binaryMatrix.Get(i, j) == 0 { break }
            first = j
            found = true // checking whether there is a 1 in the whole matrix
        }
    }
    if !found { return -1 }
    
    return first
}