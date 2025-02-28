func differenceOfSum(nums []int) int {
    var elem, digit int

    for _, num := range nums {
        elem += num
        n2 := num

        for n2 > 0 {
            digit += n2 % 10
            n2 /= 10
        }
    }

    return int(math.Abs(float64(digit - elem)))
}