func leftRigthDifference(nums []int) []int {
    prefix := []int{0}
    answer := []int{}

    for _, num := range nums {
        prefix = append(prefix, prefix[len(prefix) - 1] + num)
    }

    for i := 1; i < len(prefix); i++ {
        n := prefix[i - 1] - (prefix[len(prefix) - 1] - prefix[i])
        answer = append(answer, abs(n))
    }

    return answer
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}