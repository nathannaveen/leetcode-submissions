func countEven(num int) int {
    sum := 0
    newNum := num
    
    for newNum > 0 {
        sum += newNum % 10
        newNum /= 10
    }
    if sum % 2 == 0 {
        return num / 2
    }
    return (num - 1) / 2
}