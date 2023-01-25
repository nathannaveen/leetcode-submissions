func averageValue(nums []int) int {
    cnt, sum := 0, 0

    for _, n := range nums {
        if n % 6 == 0 {
            cnt++
            sum += n
        }
    }
    
    if cnt == 0 { return 0 }
    return sum / cnt
}