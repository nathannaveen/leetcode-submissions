func fib(n int) int {
    return helper(n)
}

func helper(n int) int {
    if n == 1 {
        return 1
    } else if n == 0 {
        return 0
    }
    return helper(n - 1) + helper(n - 2)
}