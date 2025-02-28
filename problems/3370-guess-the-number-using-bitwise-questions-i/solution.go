/** 
 * Definition of commonSetBits API.
 * func commonSetBits(num int) int;
 */

func findNumber() int {
    n := 1
    res := 0

    for i := 0; i < 31; i++ {
        x := commonSetBits(n)
        if x != 0 {
            res += n
        }
        n *= 2
    }

    return res
}
